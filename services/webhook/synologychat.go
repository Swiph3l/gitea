// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package webhook

import (
	"errors"

	webhook_model "code.gitea.io/gitea/models/webhook"
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/log"
	api "code.gitea.io/gitea/modules/structs"
)

type (
	// SynologyChatPayload represents
	SynologyChatPayload struct {
		SynologyChatRepository struct {
			URL string `json:"url"`
		} `json:"repository"`
	}

	// SynologyChatMeta contains the meta data for the webhook
	SynologyChatMeta struct {
		SynologyChatURL string `json:"synologychat_url"`
	}
)

// GetSynologyChatHook returns synologychat metadata
func GetSynologyChatHook(w *webhook_model.Webhook) *SynologyChatMeta {
	s := &SynologyChatMeta{}
	if err := json.Unmarshal([]byte(w.Meta), s); err != nil {
		log.Error("webhook.GetSynologyChatHook(%d): %v", w.ID, err)
	}
	return s
}

// JSONPayload Marshals the SynologyChatPayload to json
func (f *SynologyChatPayload) JSONPayload() ([]byte, error) {
	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

var _ PayloadConvertor = &SynologyChatPayload{}

// Create implements PayloadConvertor Create method
func (f *SynologyChatPayload) Create(p *api.CreatePayload) (api.Payloader, error) {
	return nil, nil
}

// Delete implements PayloadConvertor Delete method
func (f *SynologyChatPayload) Delete(p *api.DeletePayload) (api.Payloader, error) {
	return nil, nil
}

// Fork implements PayloadConvertor Fork method
func (f *SynologyChatPayload) Fork(p *api.ForkPayload) (api.Payloader, error) {
	return nil, nil
}

// Push implements PayloadConvertor Push method
func (f *SynologyChatPayload) Push(p *api.PushPayload) (api.Payloader, error) {
	return f, nil
}

// Issue implements PayloadConvertor Issue method
func (f *SynologyChatPayload) Issue(p *api.IssuePayload) (api.Payloader, error) {
	return nil, nil
}

// IssueComment implements PayloadConvertor IssueComment method
func (f *SynologyChatPayload) IssueComment(p *api.IssueCommentPayload) (api.Payloader, error) {
	return nil, nil
}

// PullRequest implements PayloadConvertor PullRequest method
func (f *SynologyChatPayload) PullRequest(p *api.PullRequestPayload) (api.Payloader, error) {
	return nil, nil
}

// Review implements PayloadConvertor Review method
func (f *SynologyChatPayload) Review(p *api.PullRequestPayload, event webhook_model.HookEventType) (api.Payloader, error) {
	return nil, nil
}

// Repository implements PayloadConvertor Repository method
func (f *SynologyChatPayload) Repository(p *api.RepositoryPayload) (api.Payloader, error) {
	return nil, nil
}

// Release implements PayloadConvertor Release method
func (f *SynologyChatPayload) Release(p *api.ReleasePayload) (api.Payloader, error) {
	return nil, nil
}

// GetSynologyChatPayload converts a synologychat webhook into a SynologyChatPayload
func GetSynologyChatPayload(p api.Payloader, event webhook_model.HookEventType, meta string) (api.Payloader, error) {
	s := new(SynologyChatPayload)

	synologychat := &SynologyChatMeta{}
	if err := json.Unmarshal([]byte(meta), &synologychat); err != nil {
		return s, errors.New("GetSynologyChatPayload meta json:" + err.Error())
	}
	s.SynologyChatRepository.URL = synologychat.WebhookURL
	return convertPayloader(s, p, event)
}
