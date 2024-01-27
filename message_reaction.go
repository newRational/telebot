package telebot

type ReactionType struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type MessageReaction struct {
	Chat        *Chat          `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *User          `json:"user"`
	ActorChat   *Chat          `json:"actor_chat"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

func New(emoji string) ReactionType {
	return ReactionType{
		Emoji: emoji,
	}
}
