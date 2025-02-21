package usecase

import (
	"butler/application/domains/pick_pack/models"
	outboundModel "butler/application/domains/services/outbound_order/models"
	"butler/constants"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/hasaki-tech/zeus/package/hrequest"
	"bitbucket.org/hasaki-tech/zeus/package/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	PICKER_ID = 1609
)

func (u *usecase) PickPackKafka(ctx context.Context, params *models.AutoPickPackRequest) error {
	order, err := u.outboundOrderSv.GetOne(ctx, &outboundModel.GetRequest{
		SalesOrderNumber: params.SalesOrderNumber,
	})
	if err != nil {
		return err
	}
	if order == nil {
		return errors.New("order not found")
	}

	// picking
	err = u.sendMessagePickPackKafka(ctx, order, "PICKING")
	if err != nil {
		return err
	}

	// picked
	err = u.sendMessagePickPackKafka(ctx, order, "PICKED")
	if err != nil {
		return err
	}

	// packing
	err = u.sendMessagePickPackKafka(ctx, order, "PACKING")
	if err != nil {
		return err
	}

	// packed
	err = u.sendMessagePickPackKafka(ctx, order, "PACKED")
	if err != nil {
		return err
	}

	// ready to ship
	err = u.sendMessagePickPackKafka(ctx, order, "READY_TO_SHIP")
	if err != nil {
		return err
	}

	// shipped
	err = u.sendMessagePickPackKafka(ctx, order, "SHIPPED")
	if err != nil {
		return err
	}

	// get link shipment
	linkShipment, err := u.getLinkShipment(ctx, order)
	if err != nil {
		return err
	}

	// scan shipper
	err = u.scanShipper(ctx, linkShipment)
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	// verify shipment
	err = u.sendMsgVerifyShipment(ctx, order, linkShipment)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) sendMessagePickPackKafka(ctx context.Context, order *outboundModel.OutboundOrder, status string) error {

	orderID, _ := strconv.ParseInt(order.SalesOrderId, 10, 64)
	payload := &models.WmsOrderPayload{
		OrderID:        orderID,
		OrderNumber:    order.SalesOrderNumber,
		OrderStatus:    status,
		Items:          []*models.WmsOrderPayloadItem{},
		OrderType:      order.OutboundOrderType,
		ShippingUnitId: order.ShippingUnitId,
	}
	if order.OutboundOrderType == constants.OUTBOUND_ORDER_TYPE_ORDER && order.OwnerId == constants.OWNER_ID_OMS {
		if orderID > 0 {
			payload.ShipmentId = orderID
			payload.OrderType = "SHIPMENT_ORDER"
		}
	}

	stockID, _ := strconv.ParseInt(order.WarehouseCode, 10, 64)
	payload.StockID = int(stockID)

	if status == "PICKING" {
		payload.PickerID = PICKER_ID
	}

	if status == "PICKED" {
		payload.PackerID = PICKER_ID
		payload.PackedAt = time.Now()
	}

	if status == "PACKED" {
		payload.PackerID = PICKER_ID
		payload.PackedAt = time.Now()
	}

	if status == "SHIPPED" {
		payload.PackerID = PICKER_ID
		payload.PackedAt = time.Now()
		payload.ShipperId = PICKER_ID
		payload.ShippedAt = time.Now()
	}

	outboundItems, err := u.outboundOrderSv.GetListOutboundItems(ctx, order.OutboundOrderId)
	if err != nil {
		return err
	}

	// ko lấy số lượng con lẻ trong combo
	outboundItems = utils.Where(outboundItems, func(i *outboundModel.OutboundOrderItem) bool {
		if i.ComboQty > 0 && i.ComboQty == i.Quantity {
			return false
		}

		return true
	})

	for _, item := range outboundItems {
		payload.Items = append(payload.Items, &models.WmsOrderPayloadItem{
			Sku: item.Sku,
		})
	}

	// send kafka
	err = u.lib.KafkaPublisherQc.WriteByKey(payload, order.SalesOrderNumber, "qc.wms.outbox.order.pick_pack")
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	return nil
}

func (u *usecase) scanShipper(ctx context.Context, trackingNo string) error {
	baseUrl := "https://test.hasakiexpress.vn/api/orders/scan-code"

	// Create form data

	form := url.Values{}
	form.Add("code", trackingNo)
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	// Set headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJpdiI6IlpxT0QyNDFaSURyRzI4bWtJYytkbFE9PSIsInZhbHVlIjoiZXc4aW8vRGc3L202TFB1dnZ3ZjJVdz09IiwibWFjIjoiZTUyYTEyZTUwYTAyMjg1YzMzZjM2MzM4NDhjMDE1YTkyY2YyYTYyMmY4NmEyZWU0Y2NhMjE1MjRjMWI4MjllOCIsInRhZyI6IiJ967593ea879479")
	req.Header.Set("mobile-deviceid", "dx6C5xXpQ-ig7fIs8aRp45:APA91bFTQLb5Uw4b8XKAskm8AvTmSH_xaXUHuWK-1A3ybGel7y8h9YjuLkIq7YW5aaA7FmoLWRQ3_nq82XJW5mzL34aL2B9HuFipIoyrD3GDZFDGEh5GLsnipO33in2OoaBLyBKQm8-j")
	req.Header.Set("mobile-version", "2.0.22")
	req.Header.Set("mobile-platform", "android")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	return nil
}

func (u *usecase) getLinkShipment(_ context.Context, order *outboundModel.OutboundOrder) (string, error) {
	url := "https://oms-gw-qc.inshasaki.com/api/v1/oms/shipment/shipping-document?shippment_id={{shipment_id}}"
	url = strings.ReplaceAll(url, "{{shipment_id}}", order.SalesOrderId)

	type ResponseApiGetLinkShipmentOms struct {
		ShippingDocumentURL string `json:"shippment_document_url"`
		TrackingNo          string `json:"tracking_no"`
	}

	statusCode, bodyResponse, err := hrequest.MakeRequest(nil, http.MethodGet,
		url,
		nil, TIMEOUT_20S, NO_RETRY, NO_DELAY, apiRetryCondition)
	if err != nil {
		return "", err
	}
	if statusCode < 200 || statusCode >= 400 {
		return "", status.Errorf(codes.Internal, "error get link shipment oms status code: %v - res: %v", statusCode, string(bodyResponse))
	}
	apiResponse := &ResponseApiGetLinkShipmentOms{}
	if err = json.Unmarshal(bodyResponse, apiResponse); err != nil {
		return "", status.Errorf(codes.Internal, "error unmarshal get link shipment oms response: %v, body: %v", err, string(bodyResponse))
	}
	return apiResponse.TrackingNo, nil
}

func (u *usecase) sendMsgVerifyShipment(_ context.Context, order *outboundModel.OutboundOrder, trackingNo string) error {

	type message struct {
		ShipperCode string `json:"shipper_code"`
		StockID     string `json:"stock_id"`
		Orders      []struct {
			Code        string `json:"code"`
			PackageCode string `json:"package_code"`
		} `json:"orders"`
		VerifiedAt      time.Time `json:"verified_at"`
		VerifiedBy      int64     `json:"verified_by"`
		VerifiedByEmail string    `json:"verified_by_email"`
		VerifiedByName  string    `json:"verified_by_name"`
	}

	payload := &message{
		ShipperCode:     "45454554",
		StockID:         order.WarehouseCode,
		VerifiedAt:      time.Now(),
		VerifiedBy:      16014,
		VerifiedByEmail: "toinv@hasaki.vn",
		VerifiedByName:  "Nguyễn Văn Tới",
		Orders: []struct {
			Code        string "json:\"code\""
			PackageCode string "json:\"package_code\""
		}{{
			Code:        trackingNo,
			PackageCode: order.SalesOrderNumber,
		},
		},
	}

	if err := u.lib.KafkaPublisherQc.WriteByKey(payload, order.WarehouseCode, "qc.wms.outbox.verify_delivery_orders"); err != nil {
		return err
	}

	return nil
}
