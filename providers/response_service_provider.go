package providers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
)

type ResponseServiceProvider struct{}

func (c ResponseServiceProvider) Register(container inter.Container) inter.Container {
	// Response decorators are responsible for modifying the response object.
	// All these decorators will be used to customize the response object.
	// Extent the decorators to further customize the response object
	container.Extend("response_decorators", func(service interface{}) interface{} {
		decorators := service.([]inter.ResponseDecorator)
		decorators = append(
			decorators,
			// place your custom decorators here
		)

		return decorators
	})

	// Outcome encoders are responsible for converting an object
	// to a string. One encoder will eventually be used.
	container.Bind("outcome_html_encoders", outcome.HtmlEncoders)
	container.Bind("outcome_json_encoders", outcome.JsonEncoders)

	return container
}
