package dto


type CreateReportRequest struct {
	ProjectID string `json:"project_id" validate:"required,uuid"`
	URLPDF    string `json:"url_pdf" validate:"required,url"` 
}

type UpdateReportRequest struct {
	ID        string `json:"id" validate:"required,uuid"`
	ProjectID string `json:"project_id" validate:"omitempty,uuid"`
	URLPDF    string `json:"url_pdf" validate:"omitempty,url"` 
}

type ReportResponse struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	URLPDF    string `json:"url_pdf"`
	GeneratedAt string `json:"generated_at"` 
}


