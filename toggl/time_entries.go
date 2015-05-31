package toggl

import (
	"fmt"
	"time"
	"sort"
)

type TimeEntriesService struct {
	client *ApiClient
}

type TimeEntry struct {
	ID 				uint 		`json:"id,omitempty"`
	WorkspaceID		uint 		`json:"wid,omitempty"`
	ProjectID		uint 		`json:"pid,omitempty"`
	Description		string		`json:"description"`
	Duration		uint		`json:"duration,omitempty"`
	Start			time.Time	`json:"start,omitempty"`
	Stop			time.Time	`json:"stop,omitempty"`
	Tags 			[]string	`json:"tags"`
	CreatedWith		string     	`json:"created_with"`
}

func (timeEntry *TimeEntry) IsPersisted() bool {
	return timeEntry.ID > 0
}

type TimeEntryRequest struct {
	TimeEntry		TimeEntry	`json:"time_entry"`
}

type TimeEntryResponse struct {
	TimeEntry		TimeEntry	`json:"data"`
}

type byStopTime []TimeEntry
func (v byStopTime) Len() int { return len(v) }
func (v byStopTime) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v byStopTime) Less(i, j int) bool { return v[i].Stop.Unix() > v[j].Stop.Unix() }

func (service *TimeEntriesService) All() *[]TimeEntry {
	entries := new([]TimeEntry)

	service.client.DoRequest("GET", "/time_entries", nil, entries)

	sort.Sort(byStopTime(*entries))

	return entries
}

func (service *TimeEntriesService) GetByID(id uint) TimeEntry {
	response := new(TimeEntryResponse)

	service.client.DoRequest("GET", fmt.Sprintf("/time_entries/%d", id), nil, response)

	return response.TimeEntry
}

func (service *TimeEntriesService) Current() TimeEntry {
	response := new(TimeEntryResponse)

	service.client.DoRequest("GET", "/time_entries/current", nil, response)

	return response.TimeEntry
}

func (service *TimeEntriesService) Create(timeEntry TimeEntry) TimeEntry {
	if len(timeEntry.CreatedWith) == 0 {
		timeEntry.CreatedWith = "go-toggl client"
	}

	request 	:= TimeEntryRequest{TimeEntry: timeEntry}
	response	:= new(TimeEntryResponse)

	service.client.DoRequest("POST", "/time_entries", request, response)

	return response.TimeEntry
}

func (service *TimeEntriesService) Update(timeEntry TimeEntry) TimeEntry {
	request		:= TimeEntryRequest{TimeEntry: timeEntry}
	response 	:= new(TimeEntryResponse)

	service.client.DoRequest("PUT", fmt.Sprintf("/time_entries/%d", timeEntry.ID), request, response)

	return response.TimeEntry
}

func (service *TimeEntriesService) Delete(timeEntry TimeEntry) {
	request	:= TimeEntryRequest{TimeEntry: timeEntry}

	service.client.DoRequest("DELETE", fmt.Sprintf("/time_entries/%d", timeEntry.ID), request, nil)
}

func (service *TimeEntriesService) Start(timeEntry TimeEntry) TimeEntry {
	timeEntry.Stop = time.Time{}

	if len(timeEntry.CreatedWith) == 0 {
		timeEntry.CreatedWith = "go-toggl client"
	}

	request 	:= TimeEntryRequest{TimeEntry: timeEntry}
	response	:= new(TimeEntryResponse)

	service.client.DoRequest("POST", "/time_entries/start", request, response)

	return response.TimeEntry
}

func (service *TimeEntriesService) Stop(timeEntry TimeEntry) TimeEntry {
	response	:= new(TimeEntryResponse)

	service.client.DoRequest("PUT", fmt.Sprintf("/time_entries/%d/stop", timeEntry.ID), nil, response)

	return response.TimeEntry
}