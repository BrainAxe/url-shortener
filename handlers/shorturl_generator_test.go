package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const SecretKey = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.ussc.gov/sites/default/files/pdf/training/annual-national-training-seminar/2018/Emerging_Tech_Bitcoin_Crypto.pdf"
	shortLink_1 := GenerateShortLink(initialLink_1, SecretKey)

	initialLink_2 := "https://community.openai.com/t/foundational-must-read-gpt-llm-papers/197003"
	shortLink_2 := GenerateShortLink(initialLink_2, SecretKey)

	initialLink_3 := "https://arxiv.org/abs/2403.03883?trk=article-ssr-frontend-pulse_little-text-block"
	shortLink_3 := GenerateShortLink(initialLink_3, SecretKey)

	assert.Equal(t, shortLink_1, "iGpxotoX")
	assert.Equal(t, shortLink_2, "cu7WUX8c")
	assert.Equal(t, shortLink_3, "Bc6G6hVe")
}
