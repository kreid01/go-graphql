package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"strconv"

	"kreid.com/graphl-go/graph/model"
)

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	message := &model.Message{
		User:    input.User,
		Date:    input.Date,
		Content: input.Content,
	}

	count, err := r.DB.Model(message).Count()
	if err != nil {

		return nil, fmt.Errorf("failed to count existing messages: %v", err)
	}

	message.ID = strconv.Itoa(count + 1)

	_, err = r.DB.Model(message).Insert()
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %v", err)
	}

	r.ChatMessages = append(r.ChatMessages, message)

	for _, observer := range r.ChatObservers {
		observer <- r.ChatMessages
	}

	return message, nil
}

// PostChannel is the resolver for the postChannel field.
func (r *mutationResolver) PostChannel(ctx context.Context, input model.ChannelInput) (*model.Channel, error) {
	panic(fmt.Errorf("not implemented: PostChannel - postChannel"))
}

// Message is the resolver for the message field.
func (r *queryResolver) Message(ctx context.Context, id string) (*model.Message, error) {
	message := &model.Message{ID: id}

	err := r.DB.Model(message).WherePK().Select()
	if err != nil {
		return nil, fmt.Errorf("not found")
	}

	return message, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	var messages []*model.Message

	err := r.DB.Model(&messages).Select()
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// ChatMessage is the resolver for the chatMessage field.
func (r *subscriptionResolver) ChatMessage(ctx context.Context) (<-chan []*model.Message, error) {
	panic(fmt.Errorf("not implemented: ChatMessage - chatMessage"))
}

// Channel is the resolver for the channel field.
func (r *subscriptionResolver) Channel(ctx context.Context, id string) (<-chan *model.Channel, error) {
	panic(fmt.Errorf("not implemented: Channel - channel"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
