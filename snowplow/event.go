package snowplow

// Event represents an individual atomic event. This is a mixture of what is
// needed to decode JSON messages from clients as well as what columns in the
// database these are serialized to
type Event struct {

	// Application
	AppID    string `json:"aid,omitempty" db:"app_id"`
	Platform string `json:"p,omitempty" db:"platform"`

	// Date/Time
	ETLTimestamp       string `db:"etl_tstamp"`
	CollectorTimestamp string `'db:"collector_tstamp"`
	DeviceTimestamp    string `json:"dtm,omitempty" db:"dvce_tstamp"`

	// Event
	Event         string `json:"e,omitempty" db:"event"`
	EventID       string `json:"eid,omitempty" db:"event_id"`
	TransactionID string `json:"tid,omitempty" db:"txn_id"`

	// Namespaceing and versioning
	TrackerName      string `db:"name_tracker"`
	TrackerVersion   string `json:"tv,omitempty" db:"v_tracker"`
	CollectorVersion string `db:"v_collector"`
	ETLVersion       string `db:"v_etl"`

	// User and visit
	UserID           string `json:"uid,omitempty" db:"user_id"`
	UserIPAddress    string `json:"ip,omitempty" db:"user_ipaddress"`
	UserFingerprint  string `db:"user_fingerprint"`
	DomainUserID     string `json:"duid,omitempty" db:"domain_userid"`
	DomainSessionIDX int16  `db:"domain_sessionidx"`
	NetworkUserID    string `json:"tnuid,omitempty" db:"domain_userid"`

	// Location
	GeoCountry    string  `db:"geo_country"`
	GeoRegion     string  `db:"geo_region"`
	GeoCity       string  `db:"geo_city"`
	GeoZipcode    string  `db:"geo_zipcode"`
	GeoLatitude   float32 `db:"geo_latitude"`
	GeoLongitude  float32 `db:"geo_longitude"`
	GeoRegionName string  `db:"geo_region_name"`

	// IP Lookups
	IPISP          string `db:"ip_isp"`
	IPOrganization string `db:"ip_organization"`
	IPDomain       string `db:"ip_domain"`
	IPNetspeed     string `db:"ip_netspeed"`

	// Page
	PageURL      string `json:"url,omitempty" db:"page_url"`
	PageTitle    string `json:"page,omitempty" db:"page_title"`
	PageReferrer string `json:"refr,omitempty" db:"page_referrer"`

	// Page URL Components
	PageURLScheme   string `db:"page_urlscheme"`
	PageURLHost     string `db:"page_urlhost"`
	PageURLPort     int32  `db:"page_urlport"`
	PageURLPath     string `db:"page_urlpath"`
	PageURLQuery    string `db:"page_urlquery"`
	PageURLFragment string `db:"page_urlfragment"`

	// Referrer URL Components
	ReferrerURLScheme   string `db:"refr_urlscheme"`
	ReferrerURLHost     string `db:"refr_urlhost"`
	ReferrerURLPort     int32  `db:"refr_urlport"`
	ReferrerURLPath     string `db:"refr_urlpath"`
	ReferrerURLQuery    string `db:"refr_urlquery"`
	ReferrerURLFragment string `db:"refr_urlfragment"`

	// Referrer Details
	ReferrerMedium string `db:"refr_medium"`
	ReferrerSource string `db:"refr_source"`
	ReferrerTerm   string `db:"refr_term"`

	// Marketing
	MarketingMedium   string `db:"mkt_medium"`
	MarketingSource   string `db:"mkt_source"`
	MarketingTerm     string `db:"mkt_term"`
	MarketingContent  string `db:"mkt_content"`
	MarketingCampaign string `db:"mkt_campaign"`

	// Custom Contexts
	Contexts        string `json:"co,omitempty" db:"contexts"`
	ContextsEncoded string `json:"cx,omitempty"`

	// Custom Structured Event
	StructuredEventCategory string `json:"se_ca,omitempty" db:"se_category"`
	StructuredEventAction   string `json:"se_ac,omitempty" db:"se_action"`
	StructuredEventLabel    string `json:"se_la,omitempty" db:"se_label"`
	StructuredEventProperty string `json:"se_pr,omitempty" db:"se_property"`
	StructuredEventValue    string `json:"se_va,omitempty" db:"se_value"`

	// Unstructured Event
	UnstructuredEvent        string `json:"ue_pr,omitempty" db:"unstruct_event"`
	UnstructuredEventEncoded string `json:"ue_px,omitempty"`

	// Ecommerce
	TransactionOrderID      string `json:"tr_id,omitempty" db:"tr_orderid"`
	TransactionAffiliation  string `json:"tr_af,omitempty" db:"tr_affiliation"`
	TransactionTotal        string `json:"tr_tt,omitempty" db:"tr_total"`
	TransactionTax          string `json:"tr_tx,omitempty" db:"tr_tax"`
	TransactionShipping     string `json:"tr_sh,omitempty" db:"tr_shipping"`
	TransactionCity         string `json:"tr_ci,omitempty" db:"tr_city"`
	TransactionState        string `json:"tr_st,omitempty" db:"tr_state"`
	TransactionCountry      string `json:"tr_co,omitempty" db:"tr_country"`
	TransactionItemID       string `json:"ti_id,omitempty" db:"ti_orderid"`
	TransactionItemSKU      string `json:"ti_sk,omitempty" db:"ti_sku"`
	TransactionItemName     string `json:"ti_nm,omitempty" db:"ti_name"`
	TransactionItemCategory string `json:"ti_ca,omitempty" db:"ti_category"`
	TransactionItemPrice    string `json:"ti_pr,omitempty" db:"ti_price"`
	TransactionItemQuantity string `json:"ti_qu,omitempty" db:"ti_quantity"`

	// Page Ping
	PPXOffsetMin int32 `json:"pp_mix,omitempty" db:"pp_xoffset_min"`
	PPXOffsetMax int32 `json:"pp_max,omitempty" db:"pp_xoffset_max"`
	PPYOffsetMin int32 `json:"pp_miy,omitempty" db:"pp_yoffset_min"`
	PPYOffsetMax int32 `json:"pp_may,omitempty"db:"pp_yoffset_max"`

	// User Agent
	UserAgent string `json:"ua,omitempty" db:"useragent"`

	// Browser
	BrName           string `db:"br_name"`
	BrFamily         string `db:"br_family"`
	BrVersion        string `db:"br_version"`
	BrType           string `db:"br_type"`
	BrRenderer       string `db:"br_renderengine"`
	BrLangauge       string `db:"br_lang"`
	BrFeatPDF        bool   `json:"f_pdf,omitempty" db:"br_features_pdf"`
	BrFeatFl         bool   `json:"f_fla,omitempty" db:"br_features_flash"`
	BrFeatJava       bool   `json:"f_java,omitempty" db:"br_features_java"`
	BrFeatDir        bool   `json:"f_dir,omitempty" db:"br_features_director"`
	BrFeatQT         bool   `json:"f_qt,omitempty" db:"br_features_quicktime"`
	BrFeatRealPlayer bool   `json:"f_realp,omitempty" db:"br_features_realplayer"`
	BrFeatWinMedia   bool   `json:"f_wma,omitempty" db:"br_features_windowsmedia"`
	BrFeatGears      bool   `json:"f_gears,omitempty"db:"br_features_gears"`
	BrFeatSilver     bool   `db:"br_features_silverlight"`
	BrCookies        bool   `db:"br_cookies"`
	BrColorDepth     string `db:"br_colordepth"`
	BrViewWidth      int32  `db:"br_viewwidth"`
	BrViewHeight     int32  `db:"br_viewheight"`

	// Operating System
	OSName         string `db:"os_name"`
	OSFamily       string `db:"os_family"`
	OSManufacturer string `db:"os_manufacturer"`
	OSTimeZone     string `db:"os_timezone"`

	// Device/Hardware
	DeviceType         string `db:"dvce_type"`
	DeviceIsMobile     bool   `db:"dvce_ismobile"`
	DeviceScreenWidth  int32  `db:"dvce_screenwidth"`
	DeviceStringHeight int32  `db:"dvce_screenheight"`

	// Document
	DocCharset string `db:"doc_charset"`
	DocWidth   int32  `db:"doc_width"`
	DocHeight  int32  `db:"doc_height"`

	// Currency
	TransactionCurrency     string `json:"tr_cu,omitempty" db:"tr_currency"`
	TransactionTotalBase    string `db:"tr_total_base"`
	TransactionTaxBase      string `db:"tr_tax_base"`
	TransactionShippingBase string `db:"tr_shipping_base"`
	TransactionItemCurrency string `json:"ti_cu,omitempty" db:"ti_currency"`
	TransactionItemBase     string `db:"ti_price_base"`
	BaseCurrency            string `db:"base_currency"`

	// Geolocation
	GeoTimeZone string `db:"geo_timezone"`

	// Click ID
	MarketClickID string `db:"mkt_clickid"`
	MarketNetwork string `db:"mkt_network"`

	// ETL Tags
	ETLTags string `db:"etl_tags"`

	// Time event was sent
	DeviceSentTimestamp string `db:"dvce_sent_tstamp"`

	// Referer
	ReferrerDomainUserID    string `db:"refr_domain_userid"`
	ReferrerDeviceTimestamp string `db:"refr_dvce_timestamp"`

	// Contexts
	DerivedContexts string `db:"derived_contexts"`

	// SessionID
	DomainSessionID string `db:"domain_sessionid"`

	// Derived Timestamp
	DerivedTimestamp string `db:"derived_tstamp"`

	// JSON ONLY PROPERTIES ON INCOMING EVENT

	Namespace  string `json:"tna,omitempty"`
	Timezone   string `json:"tz,omitempty"`
	Resolution string `json:"res,omitempty"`
	Language   string `json:"lang,omitempty"`
	ColorDepth string `json:"cd,omitempty"`
	Viewport   string `json:"vp,omitempty"`
}