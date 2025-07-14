package handlers

import (
	"regexp"
	"strings"
)

type InputHandler struct{}

func NewInputHandler() *InputHandler {
    return &InputHandler{}
}

func (h *InputHandler) ParseAndNormalize(input string) string {
    normalized := strings.TrimSpace(strings.ToLower(input))
    
    reg := regexp.MustCompile(`[^\w\s]`)
    normalized = reg.ReplaceAllString(normalized, "")
    
    spaceReg := regexp.MustCompile(`\s+`)
    normalized = spaceReg.ReplaceAllString(normalized, " ")
    
    return normalized
}

func (h *InputHandler) MatchKeywords(input string, keywords []string) bool {
    normalized := h.ParseAndNormalize(input)
    
    for _, keyword := range keywords {
        if strings.Contains(normalized, strings.ToLower(keyword)) {
            return true
        }
    }
    return false
}

func (h *InputHandler) MatchPhrase(input string, phrase string) bool {
    normalized := h.ParseAndNormalize(input)
    normalizedPhrase := strings.ToLower(strings.TrimSpace(phrase))
    
    return strings.Contains(normalized, normalizedPhrase)
}

func (h *InputHandler) SimpleTokenizer(input string) []string {
    normalized := h.ParseAndNormalize(input)
    tokens := strings.Fields(normalized)
    return tokens
}

func (h *InputHandler) GetFirstWord(input string) string {
    tokens := h.SimpleTokenizer(input)
    if len(tokens) > 0 {
        return tokens[0]
    }
    return ""
}

func (h *InputHandler) GetCommand(input string) string {
    return h.GetFirstWord(input)
}