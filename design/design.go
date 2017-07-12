package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const TRAIT = "trait"

var _ = API("Sample", func() {
	Title("Sample")
	Description("Sample")
	Host("localhost:8080")
	Scheme("http")
	BasePath("/")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		Headers("Content-Type", "X-Authorization")
		MaxAge(600)
		Credentials()
	})
	Trait(TRAIT, func() {
		Response(Unauthorized, ErrorMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})

var SampleMedia = MediaType("application/vnd.sample+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("integerType", Integer, "数字（1〜10）", func() {
			Example(5)
		})
		Attribute("stringType", String, "文字（1~10文字）", func() {
			Example("あいうえお")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("example@gmail.com")
		})
		Attribute("enumType", String, "列挙型", func() {
			Example("A")
		})
		Attribute("defaultType", String, "デフォルト値", func() {
			Example("でふぉ")
		})
		Attribute("reg", String, "デフォルト値", func() {
			Example("12abc")
		})
	})
	Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
	View("default", func() {
		Attribute("id")
		Attribute("integerType")
		Attribute("stringType")
		Attribute("email")
		Attribute("enumType")
		Attribute("defaultType")
		Attribute("reg")
		Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
	})
})

var _ = Resource("Sample", func() {
	BasePath("/")
	DefaultMedia(SampleMedia)
	Action("GET", func() {
		Description("Sample")
		Routing(GET(""))
		Params(func() {
			Param("id")
			Param("integerType")
			Param("stringType")
			Param("email")
			Param("enumType")
			Param("defaultType")
			Param("reg")
			Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
		})
		Response(OK, CollectionOf(SampleMedia))
		UseTrait(TRAIT)
	})
	Action("POST", func() {
		Description("Sample")
		Routing(POST("/:id"))
		Params(func() {
			Param("id", Integer, "id")
			Required("id")
		})
		Payload(func() {
			Param("integerType")
			Param("stringType")
			Param("email")
			Param("enumType")
			Param("defaultType")
			Param("reg")
			Required("integerType", "stringType", "email", "enumType", "defaultType", "reg")
		})
		Response(OK)
		UseTrait(TRAIT)
	})
})
