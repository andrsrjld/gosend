package gosend

type CalculatePriceRequest struct {
	OriginLatLong      string `json:"origin" validate:"required"`
	DestinationLatLong string `json:"destination" validate:"required"`
	PaymentType        string `json:"payment_type" validate:"required"`
}

type PriceAndDistanceEstimate struct {
	Instant Instant `json:"Instant"`
	SameDay SameDay `json:"SameDay"`
}

type Instant struct {
	ShipmentMethod            string  `json:"shipment_method"`
	ShipmentMethodDescription string  `json:"shipment_method_description"`
	Serviceable               bool    `json:"serviceable"`
	Active                    bool    `json:"active"`
	Distance                  float64 `json:"distance"`
	RoutePolyline             string  `json:"Route_polyline"`
	Price                     Price   `json:"price"`
}

type SameDay struct {
	ShipmentMethod            string  `json:"shipment_method"`
	ShipmentMethodDescription string  `json:"shipment_method_description"`
	Serviceable               bool    `json:"serviceable"`
	Active                    bool    `json:"active"`
	Distance                  float64 `json:"distance"`
	RoutePolyline             string  `json:"Route_polyline"`
	Price                     Price   `json:"price"`
}

type Price struct {
	TotalPrice                 int64  `json:"total_price"`
	GoPayTotalPrice            int64  `json:"go_pay_total_price"`
	GoPayDiscount              int64  `json:"go_pay_discount"`
	GoPayMessage               string `json:"go_pay_message"`
	VoucherDiscount            int64  `json:"voucher_discount"`
	VoucherGoPayDiscount       int64  `json:"voucher_go_pay_discount"`
	GoPayDiscountWithVoucher   int64  `json:"go_pay_discount_with_voucher"`
	VoucherMessage             string `json:"voucher_message"`
	VoucherMonetaryValue       int64  `json:"voucher_monetary_value"`
	TotalPriceWithVoucher      int64  `json:"total_price_with_voucher"`
	GoPayTotalPriceWithVoucher int64  `json:"go_pay_total_price_with_voucher"`
}

type BookingRequest struct {
	PaymentType        int64    `json:"paymentType"`
	DeviceToken        string   `json:"deviceToken"`
	CollectionLocation string   `json:"collection_location"`
	ShipmentMethod     string   `json:"shipment_method"`
	Routes             []Routes `json:"routes"`
}

type Routes struct {
	OriginName              string           `json:"originName"`
	OriginNote              string           `json:"originNote"`
	OriginContactName       string           `json:"originContactName"`
	OriginContactPhone      string           `json:"originContactPhone"`
	OriginLatLong           string           `json:"originLatLong"`
	OriginAddress           string           `json:"originAddress"`
	DestinationName         string           `json:"destinationName"`
	DestinationNote         string           `json:"destinationNote"`
	DestionationContactName string           `json:"destinationContactName"`
	DestinationContactPhone string           `json:"destinationContactPhone"`
	DestinationLatLong      string           `json:"destinationLatLong"`
	DestinationAddress      string           `json:"destinationAddress"`
	Item                    string           `json:"item"`
	StoreOrderID            string           `json:"storeOrderId"`
	InsuranceDetails        InsuranceDetails `json:"insuranceDetails"`
}

type InsuranceDetails struct {
	Applied            string `json:"applied"`
	Fee                string `json:"fee"`
	ProductDescription string `json:"product_description"`
	ProductPrice       string `json:"product_price"`
}

type BookingResponse struct {
	ID             int64  `json:"id"`
	OrderNo        string `json:"orderNo"`
	BookingType    string `json:"bookingType"`
	StoreOrderID   string `json:"storeOrderId"`
	ErrorMessage   string `json:"errorMessage"`
	PreBook        bool   `json:"prebook"`
	PreBookMessage string `json:"prebookMessage"`
}

type BookingStatusResponse struct {
	ID                  int64            `json:"id"`
	OrderNo             string           `json:"orderNo"`
	Status              string           `json:"status"`
	DriverID            int64            `json:"driverId"`
	DriverName          string           `json:"driverName"`
	DriverPhone         string           `json:"driverPhone"`
	DriverPhoto         string           `json:"driverPhoto"`
	VehicleNumber       string           `json:"vehicleNumber"`
	TotalDiscount       int64            `json:"totalDiscount"`
	TotalPrice          int64            `json:"totalPrice"`
	ReceiverName        string           `json:"receiverName"`
	OrderCreatedTime    string           `json:"orderCreatedTime"`
	OrderDispatchTime   string           `json:"orderDispatchTime"`
	OrderArrivalTime    string           `json:"orderArrivalTime"`
	OrderClosedTime     string           `json:"orderClosedTime"`
	SellerAddressName   string           `json:"sellerAddressName"`
	SellerAddressDetail string           `json:"sellerAddressDetail"`
	BuyerAddressName    string           `json:"buyerAddressName"`
	BuyerAddressDetail  string           `json:"buyerAddressDetail"`
	BookingType         string           `json:"bookingType"`
	CancelDescription   string           `json:"cancelDescription"`
	StoreOrderID        string           `json:"storeOrderId"`
	LiveTrackingURL     string           `json:"liveTrackingUrl"`
	InsuranceDetails    InsuranceDetails `json:"insuranceDetails"`
}

type BookingCancelRequest struct {
	OrderNumber string `json:"orderNo"`
}
type BookingCancelResponse struct {
	StatusCode int64  `json:"statusCode"`
	Message    string `json:"message"`
}
