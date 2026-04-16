package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	recipient  string
	message    string
	mediaPath  string
	caption    string
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a WhatsApp message or media",
	Long: `Send a WhatsApp message or media file to a recipient.

Examples:
  wacli send --to 491234567890 --message "Hello World"
  wacli send --to 491234567890 --media /path/to/image.jpg --caption "Check this out"
`,
	RunE: runSend,
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&recipient, "to", "t", "", "Recipient phone number in international format (required)")
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "Text message to send")
	sendCmd.Flags().StringVarP(&mediaPath, "media", "f", "", "Path to media file to send (image, video, document)")
	sendCmd.Flags().StringVarP(&caption, "caption", "c", "", "Caption for media message")

	_ = sendCmd.MarkFlagRequired("to")
}

func runSend(cmd *cobra.Command, args []string) error {
	// Validate that at least one of message or media is provided
	if message == "" && mediaPath == "" {
		return fmt.Errorf("either --message or --media must be provided")
	}

	// Normalize the recipient phone number (remove spaces, dashes, leading +)
	recipient = normalizePhoneNumber(recipient)

	if mediaPath != "" {
		return sendMedia(recipient, mediaPath, caption)
	}

	return sendTextMessage(recipient, message)
}

// normalizePhoneNumber strips common formatting characters from a phone number.
// Also handles parentheses, e.g. (49) 123 456-7890 -> 491234567890
func normalizePhoneNumber(phone string) string {
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	phone = strings.TrimPrefix(phone, "+")
	return phone
}

// sendTextMessage sends a plain text message via the WhatsApp API.
func sendTextMessage(to, text string) error {
	apiToken := os.Getenv("WA_API_TOKEN")
	if apiToken == "" {
		return fmt.Errorf("WA_API_TOKEN environment variable is not set")
	}

	phoneNumberID := os.Getenv("WA_PHONE_NUMBER_ID")
	if phoneNumberID == "" {
		return fmt.Errorf("WA_PHONE_NUMBER_ID environment variable is not set")
	}

	client := newWhatsAppClient(apiToken, phoneNumberID)
	if err := client.SendText(to, text); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Printf("✓ Message sent to %s\n", to)
	return nil
}

// sendMedia sends a media file (image, video, or document) via the WhatsApp API.
func sendMedia(to, filePath, caption string) error {
	apiToken := os.Getenv("WA_API_TOKEN")
	if apiToken == "" {
		return fmt.Errorf("WA_API_TOKEN environment variable is not set")
	}

	phoneNumberID := os.Getenv("WA_PHONE_NUMBER_ID")
	if phoneNumberID == "" {
		return fmt.Errorf("WA_PHONE_NUMBER_ID environment variable is not set")
	}

	// Verify the file exists before attempting to send
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("media file not found: %s", filePath)
	}

	client := newWhatsAppClient(api