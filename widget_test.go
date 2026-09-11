package bigcommerce

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestWidgetResources(t *testing.T) {
	requests := 0
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		switch requests {
		case 1:
			if r.Method != http.MethodGet || r.URL.Path != "/v3/content/widget-templates" {
				t.Fatalf("unexpected template list request: %s %s", r.Method, r.URL.Path)
			}
			if r.URL.Query().Get("widget_template_kind") != "custom" || r.URL.Query().Get("limit") != "25" {
				t.Fatalf("unexpected template query: %s", r.URL.RawQuery)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": []WidgetTemplate{{UUID: "template-id", Name: "Hero"}},
				"meta": MetaData{Pagination: Pagination{Total: 1}},
			})
		case 2:
			if r.Method != http.MethodPost || r.URL.Path != "/v3/content/widgets" {
				t.Fatalf("unexpected widget create request: %s %s", r.Method, r.URL.Path)
			}
			var body CreateWidgetParams
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body.WidgetTemplateUUID != "template-id" || body.WidgetConfiguration["title"] != "Welcome" {
				t.Fatalf("unexpected widget body: %#v", body)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": Widget{UUID: "widget-id", Name: body.Name}})
		case 3:
			if r.Method != http.MethodPut || r.URL.Path != "/v3/content/widgets/widget-id" {
				t.Fatalf("unexpected widget update request: %s %s", r.Method, r.URL.Path)
			}
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(body, map[string]any{"description": "", "upgrade": false}) {
				t.Fatalf("update did not preserve explicit zero values: %#v", body)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": Widget{UUID: "widget-id"}})
		case 4:
			if r.Method != http.MethodPost || r.URL.Path != "/v3/content/placements" {
				t.Fatalf("unexpected placement request: %s %s", r.Method, r.URL.Path)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": Placement{UUID: "placement-id"}})
		case 5:
			if r.Method != http.MethodGet || r.URL.Path != "/v3/content/regions" || r.URL.Query().Get("template_file") != "pages/home" {
				t.Fatalf("unexpected regions request: %s %s?%s", r.Method, r.URL.Path, r.URL.RawQuery)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": []ThemeRegion{{Name: "home_below_carousel"}}})
		case 6:
			if r.Method != http.MethodDelete || r.URL.Path != "/v3/content/placements/placement-id" {
				t.Fatalf("unexpected delete request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected extra request")
		}
	}))
	defer closeServer()

	templates, meta, err := client.V3.GetWidgetTemplates(WidgetTemplateQueryParams{Limit: 25, WidgetTemplateKind: "custom"})
	if err != nil || len(templates) != 1 || meta.Pagination.Total != 1 {
		t.Fatalf("unexpected templates: %#v %#v %v", templates, meta, err)
	}
	widget, err := client.V3.CreateWidget(CreateWidgetParams{
		Name: "Hero", WidgetTemplateUUID: "template-id", WidgetConfiguration: map[string]any{"title": "Welcome"},
	})
	if err != nil || widget.UUID != "widget-id" {
		t.Fatalf("unexpected widget: %#v %v", widget, err)
	}
	empty, no := "", false
	if _, err := client.V3.UpdateWidget(widget.UUID, UpdateWidgetParams{Description: &empty, Upgrade: &no}); err != nil {
		t.Fatal(err)
	}
	placement, err := client.V3.CreatePlacement(CreatePlacementParams{
		WidgetUUID: widget.UUID, TemplateFile: "pages/home", Status: "active", Region: "home_below_carousel",
	})
	if err != nil || placement.UUID != "placement-id" {
		t.Fatalf("unexpected placement: %#v %v", placement, err)
	}
	regions, err := client.V3.GetThemeRegions(ThemeRegionQueryParams{TemplateFile: "pages/home"})
	if err != nil || len(regions) != 1 {
		t.Fatalf("unexpected regions: %#v %v", regions, err)
	}
	if err := client.V3.DeletePlacement(placement.UUID); err != nil {
		t.Fatal(err)
	}
}

func TestWidgetListCommaSeparatedFilters(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("name:in") != "Hero,Footer" || r.URL.Query().Get("channel_id:in") != "1,2" {
			t.Fatalf("unexpected filters: %s", r.URL.RawQuery)
		}
		writeJSON(t, w, http.StatusOK, map[string]any{"data": []Widget{}})
	}))
	defer closeServer()

	_, _, err := client.V3.GetWidgets(WidgetQueryParams{Names: []string{"Hero", "Footer"}, ChannelIDs: []int{1, 2}})
	if err != nil {
		t.Fatal(err)
	}
}
