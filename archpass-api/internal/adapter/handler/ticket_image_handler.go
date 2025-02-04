package handler

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
	"github.com/pragma-collective/archpass/internal/domain/dto"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"os"
)

// Helper to parse hex color
func hexToColor(hex string) color.Color {
	var r, g, b uint8
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func generateTicket(eventName, eventLocation, eventDate, attendeeName, ticketName string) (*bytes.Buffer, error) {
	const (
		width  = 741
		height = 1202
	)
	dc := gg.NewContext(width, height)
	baseImagePath := "assets/tickets/archpass-ticket-2.png"

	baseImageFile, err := os.Open(baseImagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open base ticket image: %v", err)
	}
	defer baseImageFile.Close()

	baseImage, _, err := image.Decode(baseImageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base ticket image: %v", err)
	}

	dc.DrawImage(baseImage, 0, 0)

	// Ticket Name Section
	fontPath := "assets/fonts/Space_Grotesk/SpaceGrotesk-VariableFont_wght.ttf"
	fontSize := 52
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawStringAnchored(ticketName, float64(width)/2, 40, 0.5, 1)

	// Event logo section
	logoBoxWidth := 588.0
	logoBoxHeight := 120.0

	logoBoxX := (float64(width) - logoBoxWidth) / 2
	logoBoxY := float64(height) * 0.15

	dc.DrawRoundedRectangle(logoBoxX, logoBoxY, logoBoxWidth, logoBoxHeight, 10)
	dc.SetColor(hexToColor("#0C1527"))
	dc.Fill()

	// todo(jhudiel) - implement dynamic logo here, use text placeholder for now.
	fontSize = 32
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(color.White)
	dc.DrawStringAnchored("EVENT LOGO HERE", logoBoxX+logoBoxWidth/2, logoBoxY+logoBoxHeight/2, 0.5, 0.5)

	// Event Location and Date Section
	locationY := logoBoxY + 160.0
	lineWidth := 588.0
	lineX := (float64(width) - lineWidth) / 2

	dc.SetColor(hexToColor("#5666F6"))
	dc.SetLineWidth(2)
	dc.DrawLine(lineX, locationY, lineX+lineWidth, locationY)
	dc.Stroke()

	fontPath = "assets/fonts/Space_Grotesk/SpaceGrotesk-Medium.ttf"
	fontSize = 30
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(hexToColor("#0C1527"))

	locationText := eventLocation
	dc.DrawStringAnchored(locationText, lineX+10, locationY+38, 0, 0.5)

	starText := "*"
	dc.SetColor(hexToColor("#5666F6"))
	dc.DrawStringAnchored(starText, float64(width)/2, locationY+42, -1.8, 0.5)

	// Event date
	dateText := eventDate
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawStringAnchored(dateText, lineX+lineWidth-10, locationY+38, 1, 0.5)

	// Draw bottom line
	dc.SetColor(hexToColor("#5666F6"))
	dc.DrawLine(lineX, locationY+80, lineX+lineWidth, locationY+80)
	dc.Stroke()

	// Event Name Section
	eventNameY := locationY + 225.0
	eventNameWidth := 600.0
	fontPath = "assets/fonts/Space_Grotesk/SpaceGrotesk-Bold.ttf"
	fontSize = 58

	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawStringWrapped(eventName, float64(width)/2, eventNameY, 0.5, 0.5, eventNameWidth, 1.5, gg.AlignCenter)

	// Footer Section
	footerY := float64(height) - 200.0
	fontPath = "assets/fonts/Inter/Inter-Regular.ttf"
	fontSize = 44

	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(hexToColor("#0C1527"))

	attendeeNameText := fmt.Sprintf("[ %s ]", attendeeName)
	dc.DrawStringAnchored(attendeeNameText, float64(width)/2, footerY, 0.5, 0.5)

	// todo(jhudiel) - implement barcode

	// todo(jhudiel) - Implement dynamic event url
	fontPath = "assets/fonts/Inter/Inter-Regular.ttf"
	fontSize = 32 // Smaller font size for the URL
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawStringAnchored("EVENT.COM", float64(width)/2, footerY+80.0, 0.5, 0.5)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, fmt.Errorf("failed to encode image to PNG: %v", err)
	}

	return &buf, nil
}

func generateTicket3(eventName, eventLocation, eventDate, attendeeName, ticketName string) (*bytes.Buffer, error) {
	const (
		width  = 1828
		height = 678
	)
	dc := gg.NewContext(width, height)
	baseImagePath := "assets/tickets/archpass-ticket-3.png"

	baseImageFile, err := os.Open(baseImagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open base ticket image: %v", err)
	}
	defer baseImageFile.Close()

	baseImage, _, err := image.Decode(baseImageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base ticket image: %v", err)
	}

	dc.DrawImage(baseImage, 0, 0)

	// note(jhudiel) - Adjust this value to shift the entire right section dynamically
	offsetX := 20.0

	// Event Name Section
	fontPath := "assets/fonts/Space_Grotesk/SpaceGrotesk-Bold.ttf"
	fontSize := 60

	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	dc.SetColor(hexToColor("#0C1527"))
	eventNameX := float64(width)*0.64 + offsetX
	eventNameY := float64(height) * 0.15
	eventNameWidth := 600.0

	lines := dc.WordWrap(eventName, eventNameWidth)
	lineHeight := dc.FontHeight() * 1.5
	eventNameWrappedHeight := float64(len(lines)) * lineHeight

	dc.DrawStringWrapped(eventName, eventNameX, eventNameY, 0, 0, eventNameWidth, 1.5, gg.AlignLeft)

	// Event Location and Date Section
	fontPath = "assets/fonts/Bebas_Neue/BebasNeue-Regular.ttf"
	fontSize = 40
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	dc.SetColor(hexToColor("#0C1527"))

	eventDetails := fmt.Sprintf("%s | %s", eventLocation, eventDate)
	eventDetailsX := eventNameX
	eventDetailsY := eventNameY + eventNameWrappedHeight + 20.0
	dc.DrawStringAnchored(eventDetails, eventDetailsX, eventDetailsY, 0, 0.5)

	// Attendee profile image section
	// todo(jhudiel) - dynamic avatar
	profileImagePath := fmt.Sprintf("assets/avatars/%s.png", "female_avatar")
	profileImageFile, err := os.Open(profileImagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open profile image: %v", err)
	}
	defer profileImageFile.Close()

	profileImage, _, err := image.Decode(profileImageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode profile image: %v", err)
	}

	profileX := eventDetailsX
	profileY := eventDetailsY + 80.0
	profileWidth := 150.0
	profileHeight := 150.0
	dc.DrawImageAnchored(profileImage, int(profileX+profileWidth/2), int(profileY+profileHeight/2), 0.5, 0.5)

	// Seat configuration section
	fontPath = "assets/fonts/Bebas_Neue/BebasNeue-Regular.ttf"
	fontSize = 44
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	dc.SetColor(hexToColor("#0C1527"))

	columnWidth := 120.0
	baseX := profileX + profileWidth + 80.0
	baseY := profileY

	// todo(jhudiel) - dynamic seat configuration
	dc.DrawStringAnchored("ZONE", baseX, baseY+10, 0.5, 0.5)
	fontPath = "assets/fonts/Inter/Inter-Regular.ttf"
	fontSize = 30
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.DrawStringAnchored("21", baseX, baseY+60, 0.5, 0.5)

	rowX := baseX + columnWidth
	fontPath = "assets/fonts/Bebas_Neue/BebasNeue-Regular.ttf"
	fontSize = 44
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.DrawStringAnchored("ROW", rowX, baseY+10, 0.5, 0.5)
	fontPath = "assets/fonts/Inter/Inter-Regular.ttf"
	fontSize = 30
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.DrawStringAnchored("18", rowX, baseY+60, 0.5, 0.5)

	seatX := rowX + columnWidth
	fontPath = "assets/fonts/Bebas_Neue/BebasNeue-Regular.ttf"
	fontSize = 44
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.DrawStringAnchored("SEAT #", seatX, baseY+10, 0.5, 0.5)
	fontPath = "assets/fonts/Inter/Inter-Regular.ttf"
	fontSize = 30
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.DrawStringAnchored("324", seatX, baseY+60, 0.5, 0.5)

	attendeeBoxX := baseX - 60.0
	attendeeBoxY := baseY + 100.0
	attendeeBoxWidth := columnWidth * 3
	attendeeBoxHeight := 60.0

	dc.SetLineWidth(2)
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawRoundedRectangle(attendeeBoxX, attendeeBoxY, attendeeBoxWidth, attendeeBoxHeight, 10)
	dc.Stroke()

	fontPath = "assets/fonts/Inter/Inter-Regular.ttf"
	fontSize = 36
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawStringAnchored(attendeeName, attendeeBoxX+(attendeeBoxWidth/2), attendeeBoxY+(attendeeBoxHeight/2), 0.5, 0.5)

	// Bottom Section (Barcode and Ticket Type)
	fullMidSectionWidth := profileWidth + 80.0 + attendeeBoxWidth
	barcodeWidth := fullMidSectionWidth
	barcodeHeight := 60.0
	barcodeX := profileX
	barcodeY := attendeeBoxY + 100.0

	// todo(jhudiel) - implement proper barcode, think about what to encode
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawRectangle(barcodeX, barcodeY, barcodeWidth, barcodeHeight)
	dc.Fill()

	fontPath = "assets/fonts/Bebas_Neue/BebasNeue-Regular.ttf"
	fontSize = 44
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	footerY := barcodeY + barcodeHeight
	dc.SetColor(hexToColor("#0C1527"))
	dc.DrawStringAnchored(ticketName, barcodeX+(barcodeWidth/2), footerY+30, 0.5, 0.5)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, fmt.Errorf("failed to encode image to PNG: %v", err)
	}

	return &buf, nil
}

func GenerateTicketImage(c echo.Context) error {
	var input dto.CreateTicketImageInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// todo(jhudiel) - include dynamically using the different kinds of ticket.
	imageBuffer, err := generateTicket(input.EventName, input.EventLocation, input.EventDate, input.AttendeeName, input.TicketName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	// Return the generated image as a response
	c.Response().Header().Set("Content-Type", "image/png")
	c.Response().WriteHeader(http.StatusOK)
	_, err = c.Response().Write(imageBuffer.Bytes())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Failed to write image to response"})
	}

	return nil
}
