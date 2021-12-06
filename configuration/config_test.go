package configuration_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/goodmustache/appli.ed/configuration"
)

var _ = Describe("Config", func() {
	Describe("Unmarshal", func() {
		var config *Config

		BeforeEach(func() {
			config = new(Config)
		})

		It("Unmarshals from JSON properly", func() {
			raw := `{
				"watch": [
					{
						"target_config": "/some/place/else",
						"watched_locations": [
							"/watched/location/1",
							"/watched/location/2"
						]
					}
				]
			}`
			err := json.Unmarshal([]byte(raw), config)
			Expect(err).NotTo(HaveOccurred())

			Expect(config.WatchList).To(HaveLen(1))
			Expect(config.WatchList[0].Formatter).To(BeEmpty())
			Expect(config.WatchList[0].Target).To(Equal("/some/place/else"))
			Expect(config.WatchList[0].WatchedLocations).To(
				ConsistOf(
					"/watched/location/1",
					"/watched/location/2",
				),
			)
		})
	})
})
