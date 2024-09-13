package structs

type RevenueCatEvent struct {
	Event struct {
		EventTimestampMs         int64       `json:"event_timestamp_ms"`
		ProductID                string      `json:"product_id"`
		PeriodType               string      `json:"period_type"`
		PurchasedAtMs            int64       `json:"purchased_at_ms"`
		ExpirationAtMs           int64       `json:"expiration_at_ms"`
		Environment              string      `json:"environment"`
		EntitlementID            interface{} `json:"entitlement_id"`
		EntitlementIds           []string    `json:"entitlement_ids"`
		PresentedOfferingID      interface{} `json:"presented_offering_id"`
		TransactionID            string      `json:"transaction_id"`
		OriginalTransactionID    string      `json:"original_transaction_id"`
		IsFamilyShare            bool        `json:"is_family_share"`
		CountryCode              string      `json:"country_code"`
		AppUserID                string      `json:"app_user_id"`
		Aliases                  []string    `json:"aliases"`
		OriginalAppUserID        string      `json:"original_app_user_id"`
		Currency                 string      `json:"currency"`
		Price                    float64     `json:"price"`
		PriceInPurchasedCurrency float64     `json:"price_in_purchased_currency"`
		SubscriberAttributes     struct {
			Email struct {
				UpdatedAtMs int64  `json:"updated_at_ms"`
				Value       string `json:"value"`
			} `json:"$email"`
		} `json:"subscriber_attributes"`
		Store              string      `json:"store"`
		TakehomePercentage float64     `json:"takehome_percentage"`
		OfferCode          interface{} `json:"offer_code"`
		Type               string      `json:"type"`
		ID                 string      `json:"id"`
		AppID              string      `json:"app_id"`
	} `json:"event"`
	APIVersion string `json:"api_version"`
}
