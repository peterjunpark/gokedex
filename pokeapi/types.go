package pokeapi

type LocationsRes struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

type LocationRes struct {
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name                 string `json:"name"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			Rate int `json:"rate"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			EncounterDetails []struct {
				Method struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []interface{} `json:"condition_values"`
				Chance          int           `json:"chance"`
				MaxLevel        int           `json:"max_level"`
				MinLevel        int           `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
}

type Pokemon struct {
	Sprites struct {
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackFemale       interface{} `json:"back_female"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					BackDefault      string      `json:"back_default"`
					BackShiny        string      `json:"back_shiny"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackFemale       interface{} `json:"back_female"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					BackDefault      string      `json:"back_default"`
					BackShiny        string      `json:"back_shiny"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackFemale       interface{} `json:"back_female"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					BackDefault      string      `json:"back_default"`
					BackShiny        string      `json:"back_shiny"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackFemale       interface{} `json:"back_female"`
						BackShinyFemale  interface{} `json:"back_shiny_female"`
						FrontFemale      interface{} `json:"front_female"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
						BackDefault      string      `json:"back_default"`
						BackShiny        string      `json:"back_shiny"`
						FrontDefault     string      `json:"front_default"`
						FrontShiny       string      `json:"front_shiny"`
					} `json:"animated"`
					BackFemale       interface{} `json:"back_female"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					BackDefault      string      `json:"back_default"`
					BackShiny        string      `json:"back_shiny"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontFemale      interface{} `json:"front_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontShiny       string      `json:"front_shiny"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
		Other struct {
			DreamWorld struct {
				FrontFemale  interface{} `json:"front_female"`
				FrontDefault string      `json:"front_default"`
			} `json:"dream_world"`
			Home struct {
				FrontFemale      interface{} `json:"front_female"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
				FrontDefault     string      `json:"front_default"`
				FrontShiny       string      `json:"front_shiny"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackFemale       interface{} `json:"back_female"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
				BackDefault      string      `json:"back_default"`
				BackShiny        string      `json:"back_shiny"`
				FrontDefault     string      `json:"front_default"`
				FrontShiny       string      `json:"front_shiny"`
			} `json:"showdown"`
		} `json:"other"`
		BackFemale       interface{} `json:"back_female"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		BackDefault      string      `json:"back_default"`
		BackShiny        string      `json:"back_shiny"`
		FrontDefault     string      `json:"front_default"`
		FrontShiny       string      `json:"front_shiny"`
	} `json:"sprites"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Name                   string        `json:"name"`
	PastAbilities          []interface{} `json:"past_abilities"`
	PastTypes              []interface{} `json:"past_types"`
	Types                  []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
		Slot int `json:"slot"`
	} `json:"types"`
	Stats []struct {
		Stat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
	} `json:"stats"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			Rarity int `json:"rarity"`
		} `json:"version_details"`
	} `json:"held_items"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			LevelLearnedAt int `json:"level_learned_at"`
		} `json:"version_group_details"`
	} `json:"moves"`
	GameIndices []struct {
		Version struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
		GameIndex int `json:"game_index"`
	} `json:"game_indices"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Height         int  `json:"height"`
	Order          int  `json:"order"`
	BaseExperience int  `json:"base_experience"`
	ID             int  `json:"id"`
	Weight         int  `json:"weight"`
	IsDefault      bool `json:"is_default"`
}
