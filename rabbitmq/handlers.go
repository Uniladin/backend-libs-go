package rabbitmq

import (
	"context"
	"log"
)

// HandleProfileCreatedEvent handles profile creation events
// This can be used by other services to react to profile creation
func HandleProfileCreatedEvent(ctx context.Context, body []byte) error {
	var event ProfileCreatedEvent
	if err := UnmarshalMessage(body, &event); err != nil {
		log.Printf("❌ Failed to unmarshal ProfileCreatedEvent: %v", err)
		return err
	}

	log.Printf("📥 Received ProfileCreatedEvent: UserID=%s, Username=%s, Email=%s",
		event.UserID, event.Username, event.Email)

	// TODO: Add your business logic here
	// Examples:
	// - Send welcome email
	// - Create default settings
	// - Initialize user analytics
	// - Notify other services

	return nil
}

// HandleProfileUpdatedEvent handles profile update events
func HandleProfileUpdatedEvent(ctx context.Context, body []byte) error {
	var event ProfileUpdatedEvent
	if err := UnmarshalMessage(body, &event); err != nil {
		log.Printf("❌ Failed to unmarshal ProfileUpdatedEvent: %v", err)
		return err
	}

	log.Printf("📥 Received ProfileUpdatedEvent: UserID=%s, Username=%s",
		event.UserID, event.Username)

	// TODO: Add your business logic here
	// Examples:
	// - Update cache
	// - Sync with other services
	// - Update search index

	return nil
}

// HandleProfileDeletedEvent handles profile deletion events
func HandleProfileDeletedEvent(ctx context.Context, body []byte) error {
	var event ProfileDeletedEvent
	if err := UnmarshalMessage(body, &event); err != nil {
		log.Printf("❌ Failed to unmarshal ProfileDeletedEvent: %v", err)
		return err
	}

	log.Printf("📥 Received ProfileDeletedEvent: UserID=%s", event.UserID)

	// TODO: Add your business logic here
	// Examples:
	// - Delete related data
	// - Archive user content
	// - Notify other services

	return nil
}

// HandleShareLinkCreatedEvent handles share link creation events
func HandleShareLinkCreatedEvent(ctx context.Context, body []byte) error {
	var event ShareLinkCreatedEvent
	if err := UnmarshalMessage(body, &event); err != nil {
		log.Printf("❌ Failed to unmarshal ShareLinkCreatedEvent: %v", err)
		return err
	}

	log.Printf("📥 Received ShareLinkCreatedEvent: ShareLinkID=%s, ShortCode=%s",
		event.ShareLinkID, event.ShortCode)

	// TODO: Add your business logic here
	// Examples:
	// - Update analytics
	// - Cache short code mapping
	// - Notify monitoring systems

	return nil
}

// HandleShareLinkAccessedEvent handles share link access events
func HandleShareLinkAccessedEvent(ctx context.Context, body []byte) error {
	var event ShareLinkAccessedEvent
	if err := UnmarshalMessage(body, &event); err != nil {
		log.Printf("❌ Failed to unmarshal ShareLinkAccessedEvent: %v", err)
		return err
	}

	log.Printf("📥 Received ShareLinkAccessedEvent: ShareLinkID=%s, ShortCode=%s, IP=%s",
		event.ShareLinkID, event.ShortCode, event.IPAddress)

	// TODO: Add your business logic here
	// Examples:
	// - Track analytics
	// - Update access count
	// - Detect suspicious activity
	// - Log access patterns

	return nil
}
