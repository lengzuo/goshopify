package dto

import "time"

type Webhook struct {
	Address                    string     `json:"address,omitempty"`
	APIVersion                 string     `json:"api_version,omitempty"`
	CreatedAt                  *time.Time `json:"created_at,omitempty"`
	Fields                     []string   `json:"fields,omitempty"`
	Format                     string     `json:"format,omitempty"`
	ID                         int64      `json:"id,omitempty"`
	MetafieldNamespaces        []string   `json:"metafield_namespaces,omitempty"`
	PrivateMetafieldNamespaces []string   `json:"private_metafield_namespaces,omitempty"`
	Topic                      string     `json:"topic,omitempty"`
	UpdatedAt                  *time.Time `json:"updated_at,omitempty"`
}

type WebhookCollection struct {
	Webhook Webhook `json:"webhook"`
}

type GetWebhookRequest struct {
	Address               string     `url:"address"`
	CreatedAtMax          *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin          *time.Time `url:"created_at_min,omitempty"`
	Fields                string     `url:"fields"`
	Limit                 int        `url:"limit"`
	PresentmentCurrencies string     `url:"presentment_currencies"`
	SinceID               int64      `url:"since_id"`
	Topic                 string     `url:"topic"`
	UpdatedAtMax          *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin          *time.Time `url:"updated_at_min,omitempty"`
}

type GetWebhookResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type CountWebhookRequest struct {
	Address string `url:"address"`
	Topic   string `url:"topic"`
}
