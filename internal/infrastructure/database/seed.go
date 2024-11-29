package database

import (
	"FPPBKKGo/internal/domain"
	"log"
	"math/rand"

	"gorm.io/gorm"
)

// SeedActors seeds the actors table with initial data
func SeedActors(db *gorm.DB) {
	actors := []domain.Actor{
		{Name: "Leonardo DiCaprio"},
		{Name: "Scarlett Johansson"},
		{Name: "Denzel Washington"},
	}

	if err := db.Create(&actors).Error; err != nil {
		log.Fatalf("Failed to seed actors table: %v", err)
	}
	log.Println("Actors table seeded successfully")
}

// SeedGenres seeds the genres table with initial data
func SeedGenres(db *gorm.DB) {
	genres := []domain.Genre{
		{Name: "Action", Slug: "action"},
		{Name: "Drama", Slug: "drama"},
		{Name: "Comedy", Slug: "comedy"},
		{Name: "Sci-Fi", Slug: "scifi"},
	}

	if err := db.Create(&genres).Error; err != nil {
		log.Fatalf("Failed to seed genres table: %v", err)
	}
	log.Println("Genres table seeded successfully")
}

func SeedMovies(db *gorm.DB) {
	movies := []domain.Movie{
		{
			Title: 		 "The Super Mario Bros. Movie",
			Slug: 		 "not-all-heroes-wear-capesi-some-wear-overallsi",
			Year: 		 "2023",
			Description: 	 "While working underground to fix a water main, Brooklyn plumbers—and brothers—Mario and Luigi are transported down a mysterious pipe and wander into a magical new world. But when the brothers are separated, Mario embarks on an epic quest to find Luigi.",
			ImagePath: 	 "images/movie_images/502356",
		},
		
	
		{
			Title: 		 "Barbie",
			Slug: 		 "she-s-everythingi-he-s-just-keni",
			Year: 		 "2023",
			Description: 	 "Barbie and Ken are having the time of their lives in the colorful and seemingly perfect world of Barbie Land. However, when they get a chance to go to the real world, they soon discover the joys and perils of living among humans.",
			ImagePath: 	 "images/movie_images/346698",
		},
		
	
		{
			Title: 		 "Guardians of the Galaxy Vol. 3",
			Slug: 		 "once-more-with-feelingi",
			Year: 		 "2023",
			Description: 	 "Peter Quill, still reeling from the loss of Gamora, must rally his team around him to defend the universe along with protecting one of their own. A mission that, if not completed successfully, could quite possibly lead to the end of the Guardians as we know them.",
			ImagePath: 	 "images/movie_images/447365",
		},
		
	
		{
			Title: 		 "John Wick: Chapter 4",
			Slug: 		 "no-way-backi-one-way-outi",
			Year: 		 "2023",
			Description: 	 "With the price on his head ever increasing, John Wick uncovers a path to defeating The High Table. But before he can earn his freedom, Wick must face off against a new enemy with powerful alliances across the globe and forces that turn old friends into foes.",
			ImagePath: 	 "images/movie_images/603692",
		},
		
	
		{
			Title: 		 "Spider-Man: Across the Spider-Verse",
			Slug: 		 "it-s-how-you-wear-the-mask-that-mattersi",
			Year: 		 "2023",
			Description: 	 "After reuniting with Gwen Stacy, Brooklyn’s full-time, friendly neighborhood Spider-Man is catapulted across the Multiverse, where he encounters the Spider Society, a team of Spider-People charged with protecting the Multiverse’s very existence. But when the heroes clash on how to handle a new threat, Miles finds himself pitted against the other Spiders and must set out on his own to save those he loves most.",
			ImagePath: 	 "images/movie_images/569094",
		},
		
	
		{
			Title: 		 "Ant-Man and the Wasp: Quantumania",
			Slug: 		 "witness-the-beginning-of-a-new-dynastyi",
			Year: 		 "2023",
			Description: 	 "Super-Hero partners Scott Lang and Hope van Dyne, along with with Hope's parents Janet van Dyne and Hank Pym, and Scott's daughter Cassie Lang, find themselves exploring the Quantum Realm, interacting with strange new creatures and embarking on an adventure that will push them beyond the limits of what they thought possible.",
			ImagePath: 	 "images/movie_images/640146",
		},
		
	
		{
			Title: 		 "Fast X",
			Slug: 		 "the-end-of-the-road-beginsi",
			Year: 		 "2023",
			Description: 	 "Over many missions and against impossible odds, Dom Toretto and his family have outsmarted, out-nerved and outdriven every foe in their path. Now, they confront the most lethal opponent they've ever faced: A terrifying threat emerging from the shadows of the past who's fueled by blood revenge, and who is determined to shatter this family and destroy everything—and everyone—that Dom loves, forever.",
			ImagePath: 	 "images/movie_images/385687",
		},
		
	
		{
			Title: 		 "Oppenheimer",
			Slug: 		 "the-world-forever-changesi",
			Year: 		 "2023",
			Description: 	 "The story of J. Robert Oppenheimer’s role in the development of the atomic bomb during World War II.",
			ImagePath: 	 "images/movie_images/872585",
		},
		
	
		{
			Title: 		 "Transformers: Rise of the Beasts",
			Slug: 		 "unite-or-falli",
			Year: 		 "2023",
			Description: 	 "When a new threat capable of destroying the entire planet emerges, Optimus Prime and the Autobots must team up with a powerful faction known as the Maximals. With the fate of humanity hanging in the balance, humans Noah and Elena will do whatever it takes to help the Transformers as they engage in the ultimate battle to save Earth.",
			ImagePath: 	 "images/movie_images/667538",
		},
		
	
		{
			Title: 		 "The Flash",
			Slug: 		 "worlds-collidei",
			Year: 		 "2023",
			Description: 	 "When his attempt to save his family inadvertently alters the future, Barry Allen becomes trapped in a reality in which General Zod has returned and there are no Super Heroes to turn to. In order to save the world that he is in and return to the future that he knows, Barry's only hope is to race for his life. But will making the ultimate sacrifice be enough to reset the universe?",
			ImagePath: 	 "images/movie_images/298618",
		},
		
	
		{
			Title: 		 "Elemental",
			Slug: 		 "opposites-reacti",
			Year: 		 "2023",
			Description: 	 "In a city where fire, water, land and air residents live together, a fiery young woman and a go-with-the-flow guy will discover something elemental: how much they have in common.",
			ImagePath: 	 "images/movie_images/976573",
		},
		
	
		{
			Title: 		 "Dungeons & Dragons: Honor Among Thieves",
			Slug: 		 "no-experience-necessaryi",
			Year: 		 "2023",
			Description: 	 "A charming thief and a band of unlikely adventurers undertake an epic heist to retrieve a lost relic, but things go dangerously awry when they run afoul of the wrong people.",
			ImagePath: 	 "images/movie_images/493529",
		},
		
	
		{
			Title: 		 "Evil Dead Rise",
			Slug: 		 "mommy-loves-you-to-deathi",
			Year: 		 "2023",
			Description: 	 "A reunion between two estranged sisters gets cut short by the rise of flesh-possessing demons, thrusting them into a primal battle for survival as they face the most nightmarish version of family imaginable.",
			ImagePath: 	 "images/movie_images/713704",
		},
		
	
		{
			Title: 		 "The Little Mermaid",
			Slug: 		 "watch-and-you-ll-seei-some-day-i-ll-bei-part-of-your-world%21",
			Year: 		 "2023",
			Description: 	 "The youngest of King Triton’s daughters, and the most defiant, Ariel longs to find out more about the world beyond the sea, and while visiting the surface, falls for the dashing Prince Eric. With mermaids forbidden to interact with humans, Ariel makes a deal with the evil sea witch, Ursula, which gives her a chance to experience life on land, but ultimately places her life – and her father’s crown – in jeopardy.",
			ImagePath: 	 "images/movie_images/447277",
		},
		
	
		{
			Title: 		 "Meg 2: The Trench",
			Slug: 		 "back-for-secondsi",
			Year: 		 "2023",
			Description: 	 "An exploratory dive into the deepest depths of the ocean of a daring research team spirals into chaos when a malevolent mining operation threatens their mission and forces them into a high-stakes battle for survival.",
			ImagePath: 	 "images/movie_images/615656",
		},
		
	
		{
			Title: 		 "Creed III",
			Slug: 		 "you-can-t-run-from-your-pasti",
			Year: 		 "2023",
			Description: 	 "After dominating the boxing world, Adonis Creed has thrived in his career and family life. When a childhood friend and former boxing prodigy, Damian Anderson, resurfaces after serving a long sentence in prison, he is eager to prove that he deserves his shot in the ring. The face-off between former friends is more than just a fight. To settle the score, Adonis must put his future on the line to battle Damian — a fighter with nothing to lose.",
			ImagePath: 	 "images/movie_images/677179",
		},
		
	
		{
			Title: 		 "The Pope's Exorcist",
			Slug: 		 "inspired-by-the-actual-files-of-father-gabriele-amorthi-chief-exorcist-of-the-vaticani",
			Year: 		 "2023",
			Description: 	 "Father Gabriele Amorth, Chief Exorcist of the Vatican, investigates a young boy's terrifying possession and ends up uncovering a centuries-old conspiracy the Vatican has desperately tried to keep hidden.",
			ImagePath: 	 "images/movie_images/758323",
		},
		
	
		{
			Title: 		 "Scream VI",
			Slug: 		 "new-yorki-new-rulesi",
			Year: 		 "2023",
			Description: 	 "Following the latest Ghostface killings, the four survivors leave Woodsboro behind and start a fresh chapter.",
			ImagePath: 	 "images/movie_images/934433",
		},
		
	
		{
			Title: 		 "Knock at the Cabin",
			Slug: 		 "save-your-family-or-save-humanityi-make-the-choicei",
			Year: 		 "2023",
			Description: 	 "While vacationing at a remote cabin, a young girl and her two fathers are taken hostage by four armed strangers who demand that the family make an unthinkable choice to avert the apocalypse. With limited access to the outside world, the family must decide what they believe before all is lost.",
			ImagePath: 	 "images/movie_images/631842",
		},
		
	
		{
			Title: 		 "Extraction 2",
			Slug: 		 "prepare-for-the-ride-of-your-lifei",
			Year: 		 "2023",
			Description: 	 "Tasked with extracting a family who is at the mercy of a Georgian gangster, Tyler Rake infiltrates one of the world's deadliest prisons in order to save them. But when the extraction gets hot, and the gangster dies in the heat of battle, his equally ruthless brother tracks down Rake and his team to Vienna, in order to get revenge.",
			ImagePath: 	 "images/movie_images/697843",
		},
		
	
		{
			Title: 		 "Plane",
			Slug: 		 "survive-together-or-die-alonei",
			Year: 		 "2023",
			Description: 	 "After a heroic job of successfully landing his storm-damaged aircraft in a war zone, a fearless pilot finds himself between the agendas of multiple militias planning to take the plane and its passengers hostage.",
			ImagePath: 	 "images/movie_images/646389",
		},
		
	
		{
			Title: 		 "Indiana Jones and the Dial of Destiny",
			Slug: 		 "a-legend-will-face-his-destinyi",
			Year: 		 "2023",
			Description: 	 "Finding himself in a new era, and approaching retirement, Indy wrestles with fitting into a world that seems to have outgrown him. But as the tentacles of an all-too-familiar evil return in the form of an old rival, Indy must don his hat and pick up his whip once more to make sure an ancient and powerful artifact doesn't fall into the wrong hands.",
			ImagePath: 	 "images/movie_images/335977",
		},
		
	
		{
			Title: 		 "65",
			Slug: 		 "65-million-years-agoi-prehistoric-earth-had-a-visitori",
			Year: 		 "2023",
			Description: 	 "65 million years ago, the only 2 survivors of a spaceship from Somaris that crash-landed on Earth, must fend off dinosaurs to reach the escape vessel in time before an imminent asteroid strike threatens to destroy the planet.",
			ImagePath: 	 "images/movie_images/700391",
		},
		
	
		{
			Title: 		 "Guy Ritchie's The Covenant",
			Slug: 		 "a-bondi-a-pledgei-a-commitmenti",
			Year: 		 "2023",
			Description: 	 "During the war in Afghanistan, a local interpreter risks his own life to carry an injured sergeant across miles of grueling terrain.",
			ImagePath: 	 "images/movie_images/882569",
		},
		
	
		{
			Title: 		 "Cocaine Bear",
			Slug: 		 "get-in-linei",
			Year: 		 "2023",
			Description: 	 "Inspired by a true story, an oddball group of cops, criminals, tourists and teens converge in a Georgia forest where a 500-pound black bear goes on a murderous rampage after unintentionally ingesting cocaine.",
			ImagePath: 	 "images/movie_images/804150",
		},
		
	
		{
			Title: 		 "No Hard Feelings",
			Slug: 		 "prettyi-awkwardi",
			Year: 		 "2023",
			Description: 	 "On the brink of losing her childhood home, Maddie discovers an intriguing job listing: wealthy helicopter parents looking for someone to “date” their introverted 19-year-old son, Percy, before he leaves for college. To her surprise, Maddie soon discovers the awkward Percy is no sure thing.",
			ImagePath: 	 "images/movie_images/884605",
		},
		
	
		{
			Title: 		 "Air",
			Slug: 		 "some-icons-are-meant-to-flyi",
			Year: 		 "2023",
			Description: 	 "Discover the game-changing partnership between a then undiscovered Michael Jordan and Nike's fledgling basketball division which revolutionized the world of sports and culture with the Air Jordan brand.",
			ImagePath: 	 "images/movie_images/964980",
		},
		
	
		{
			Title: 		 "Murder Mystery 2",
			Slug: 		 "this-mystery-is-deux-or-diei",
			Year: 		 "2023",
			Description: 	 "After starting their own detective agency, Nick and Audrey Spitz land a career-making case when their billionaire pal is kidnapped from his wedding.",
			ImagePath: 	 "images/movie_images/638974",
		},
		
	
		{
			Title: 		 "Ghosted",
			Slug: 		 "finding-that-special-someone-can-be-a-real-adventurei",
			Year: 		 "2023",
			Description: 	 "Salt-of-the-earth Cole falls head over heels for enigmatic Sadie — but then makes the shocking discovery that she’s a secret agent. Before they can decide on a second date, Cole and Sadie are swept away on an international adventure to save the world.",
			ImagePath: 	 "images/movie_images/868759",
		},
		
	
		{
			Title: 		 "Mission: Impossible - Dead Reckoning Part One",
			Slug: 		 "we-all-share-the-same-fatei",
			Year: 		 "2023",
			Description: 	 "Ethan Hunt and his IMF team embark on their most dangerous mission yet: To track down a terrifying new weapon that threatens all of humanity before it falls into the wrong hands. With control of the future and the world's fate at stake and dark forces from Ethan's past closing in, a deadly race around the globe begins. Confronted by a mysterious, all-powerful enemy, Ethan must consider that nothing can matter more than his mission—not even the lives of those he cares about most.",
			ImagePath: 	 "images/movie_images/575264",
		},
		
	
		{
			Title: 		 "Heart of Stone",
			Slug: 		 "defy-the-odds",
			Year: 		 "2023",
			Description: 	 "An intelligence operative for a shadowy global peacekeeping agency races to stop a hacker from stealing its most valuable — and dangerous — weapon.",
			ImagePath: 	 "images/movie_images/724209",
		},
		
	
		{
			Title: 		 "Insidious: The Red Door",
			Slug: 		 "it-ends-where-it-all-begani",
			Year: 		 "2023",
			Description: 	 "To put their demons to rest once and for all, Josh Lambert and a college-aged Dalton Lambert must go deeper into The Further than ever before, facing their family's dark past and a host of new and more horrifying terrors that lurk behind the red door.",
			ImagePath: 	 "images/movie_images/614479",
		},
		
	
		{
			Title: 		 "Operation Fortune: Ruse de Guerre",
			Slug: 		 "in-this-operationi-everyone-has-a-part-to-playi",
			Year: 		 "2023",
			Description: 	 "Special agent Orson Fortune and his team of operatives recruit one of Hollywood's biggest movie stars to help them on an undercover mission when the sale of a deadly new weapons technology threatens to disrupt the world order.",
			ImagePath: 	 "images/movie_images/739405",
		},
		
	
		{
			Title: 		 "Blue Beetle",
			Slug: 		 "jaime-reyes-is-a-superhero-whether-he-likes-it-or-noti",
			Year: 		 "2023",
			Description: 	 "Recent college grad Jaime Reyes returns home full of aspirations for his future, only to find that home is not quite as he left it. As he searches to find his purpose in the world, fate intervenes when Jaime unexpectedly finds himself in possession of an ancient relic of alien biotechnology: the Scarab.",
			ImagePath: 	 "images/movie_images/565770",
		},
		
	
		{
			Title: 		 "Renfield",
			Slug: 		 "sucks-to-be-himi",
			Year: 		 "2023",
			Description: 	 "Having grown sick and tired of his centuries as Dracula's lackey, Renfield finds a new lease on life — and maybe even redemption — when he falls for feisty, perennially angry traffic cop Rebecca Quincy.",
			ImagePath: 	 "images/movie_images/649609",
		},
		
	
		{
			Title: 		 "The Mother",
			Slug: 		 "vengeance-is-a-mother%21",
			Year: 		 "2023",
			Description: 	 "A deadly female assassin comes out of hiding to protect the daughter that she gave up years before, while on the run from dangerous men.",
			ImagePath: 	 "images/movie_images/552688",
		},
		
	
		{
			Title: 		 "Talk to Me",
			Slug: 		 "you-calli-they-ll-answeri",
			Year: 		 "2023",
			Description: 	 "When a group of friends discover how to conjure spirits using an embalmed hand, they become hooked on the new thrill, until one of them goes too far and unleashes terrifying supernatural forces.",
			ImagePath: 	 "images/movie_images/1008042",
		},
		
	
		{
			Title: 		 "Tetris",
			Slug: 		 "the-game-you-couldn-t-put-downi-the-story-you-couldn-t-make-upi",
			Year: 		 "2023",
			Description: 	 "In 1988, American video game salesman Henk Rogers discovers the video game Tetris. When he sets out to bring the game to the world, he enters a dangerous web of lies and corruption behind the Iron Curtain.",
			ImagePath: 	 "images/movie_images/726759",
		},
		
	
		{
			Title: 		 "Knights of the Zodiac",
			Slug: 		 "go-beyond-your-destinyi",
			Year: 		 "2023",
			Description: 	 "When a headstrong street orphan, Seiya, in search of his abducted sister unwittingly taps into hidden powers, he discovers he might be the only person alive who can protect a reincarnated goddess, sent to watch over humanity. Can he let his past go and embrace his destiny to become a Knight of the Zodiac?",
			ImagePath: 	 "images/movie_images/455476",
		},
		
	
		{
			Title: 		 "Red, White & Royal Blue",
			Slug: 		 "love-who-you-wanti-it-s-good-foreign-policyi",
			Year: 		 "2023",
			Description: 	 "After an altercation between Alex, the president's son, and Britain's Prince Henry at a royal event becomes tabloid fodder, their long-running feud now threatens to drive a wedge in U.S./British relations. When the rivals are forced into a staged truce, their icy relationship begins to thaw and the friction between them sparks something deeper than they ever expected.",
			ImagePath: 	 "images/movie_images/930094",
		},
		
	
		{
			Title: 		 "Hidden Strike",
			Slug: 		 "there-is-a-plani-they-just-don-t-know-what-it-isi",
			Year: 		 "2023",
			Description: 	 "Two elite soldiers must escort civilians through a gauntlet of gunfire and explosions.",
			ImagePath: 	 "images/movie_images/457332",
		},
		
	
		{
			Title: 		 "The Last Voyage of the Demeter",
			Slug: 		 "the-legend-of-dracula-is-borni",
			Year: 		 "2023",
			Description: 	 "The crew of the merchant ship Demeter attempts to survive the ocean voyage from Carpathia to London as they are stalked each night by a merciless presence onboard the ship.",
			ImagePath: 	 "images/movie_images/635910",
		},
		
	
		{
			Title: 		 "Missing",
			Slug: 		 "no-one-disappears-without-a-tracei",
			Year: 		 "2023",
			Description: 	 "When her mother disappears while on vacation in Colombia with her new boyfriend, June’s search for answers is hindered by international red tape. Stuck thousands of miles away in Los Angeles, June creatively uses all the latest technology at her fingertips to try and find her before it’s too late. But as she digs deeper, her digital sleuthing raises more questions than answers... and when June unravels secrets about her mom, she discovers that she never really knew her at all.",
			ImagePath: 	 "images/movie_images/768362",
		},
		
	
		{
			Title: 		 "Gran Turismo",
			Slug: 		 "from-gamer-to-raceri",
			Year: 		 "2023",
			Description: 	 "The ultimate wish-fulfillment tale of a teenage Gran Turismo player whose gaming skills won him a series of Nissan competitions to become an actual professional racecar driver.",
			ImagePath: 	 "images/movie_images/980489",
		},
		
	
		{
			Title: 		 "Teen Wolf: The Movie",
			Slug: 		 "the-pack-is-backi",
			Year: 		 "2023",
			Description: 	 "The wolves are howling once again, as a terrifying ancient evil emerges in Beacon Hills. Scott McCall, no longer a teenager yet still an Alpha, must gather new allies and reunite trusted friends to fight back against this powerful and deadly enemy.",
			ImagePath: 	 "images/movie_images/877703",
		},
		
	
		{
			Title: 		 "Ruby Gillman, Teenage Kraken",
			Slug: 		 "discover-the-hero-just-beneath-the-surfacei",
			Year: 		 "2023",
			Description: 	 "Ruby Gillman, a sweet and awkward high school student, discovers she's a direct descendant of the warrior kraken queens. The kraken are sworn to protect the oceans of the world against the vain, power-hungry mermaids. Destined to inherit the throne from her commanding grandmother, Ruby must use her newfound powers to protect those she loves most.",
			ImagePath: 	 "images/movie_images/1040148",
		},
		
	
		{
			Title: 		 "Teenage Mutant Ninja Turtles: Mutant Mayhem",
			Slug: 		 "heroes-aren-t-borni-they-re-mutatedi",
			Year: 		 "2023",
			Description: 	 "After years of being sheltered from the human world, the Turtle brothers set out to win the hearts of New Yorkers and be accepted as normal teenagers through heroic acts. Their new friend April O'Neil helps them take on a mysterious crime syndicate, but they soon get in over their heads when an army of mutants is unleashed upon them.",
			ImagePath: 	 "images/movie_images/614930",
		},
		
	
		{
			Title: 		 "Hypnotic",
			Slug: 		 "control-is-an-illusioni",
			Year: 		 "2023",
			Description: 	 "A detective becomes entangled in a mystery involving his missing daughter and a secret government program while investigating a string of reality-bending crimes.",
			ImagePath: 	 "images/movie_images/536437",
		},
		
	
		{
			Title: 		 "Your Place or Mine",
			Slug: 		 "two-livesi-two-citiesi-one-last-chancei",
			Year: 		 "2023",
			Description: 	 "When best friends and total opposites Debbie and Peter swap homes for a week, they get a peek into each other's lives that could open the door to love.",
			ImagePath: 	 "images/movie_images/703451",
		},
		
	
		{
			Title: 		 "Kandahar",
			Slug: 		 "the-only-thing-more-dangerous-than-the-mission-is-the-escapei",
			Year: 		 "2023",
			Description: 	 "After his mission is exposed, an undercover CIA operative stuck deep in hostile territory in Afghanistan must fight his way out, alongside his Afghan translator, to an extraction point in Kandahar, all whilst avoiding elite enemy forces and foreign spies tasked with hunting them down.",
			ImagePath: 	 "images/movie_images/717930",
		},
		
	
		{
			Title: 		 "You People",
			Slug: 		 "opposites-attracti-families-don-ti",
			Year: 		 "2023",
			Description: 	 "A new couple and their families reckon with modern love amid culture clashes, societal expectations and generational differences.",
			ImagePath: 	 "images/movie_images/866413",
		},
		
	
		{
			Title: 		 "Peter Pan & Wendy",
			Slug: 		 "escape-to-neverlandi",
			Year: 		 "2023",
			Description: 	 "Wendy Darling, a young girl afraid to leave her childhood home behind, meets Peter Pan, a boy who refuses to grow up. Alongside her brothers and a tiny fairy, Tinker Bell, she travels with Peter to the magical world of Neverland. There, she encounters an evil pirate captain, Captain Hook, and embarks on a thrilling adventure that will change her life forever.",
			ImagePath: 	 "images/movie_images/420808",
		},
		
	
		{
			Title: 		 "Beau Is Afraid",
			Slug: 		 "from-his-darkest-fears-comes-the-greatest-adventurei",
			Year: 		 "2023",
			Description: 	 "A paranoid man embarks on an epic odyssey to get home to his mother.",
			ImagePath: 	 "images/movie_images/798286",
		},
		
	
		{
			Title: 		 "A Haunting in Venice",
			Slug: 		 "death-was-only-the-beginningi",
			Year: 		 "2023",
			Description: 	 "Celebrated sleuth Hercule Poirot, now retired and living in self-imposed exile in Venice, reluctantly attends a Halloween séance at a decaying, haunted palazzo. When one of the guests is murdered, the detective is thrust into a sinister world of shadows and secrets.",
			ImagePath: 	 "images/movie_images/945729",
		},
		
	
		{
			Title: 		 "The Black Demon",
			Slug: 		 "nature-bites-backi",
			Year: 		 "2023",
			Description: 	 "Oilman Paul Sturges' idyllic family vacation turns into a nightmare when they encounter a ferocious megalodon shark that will stop at nothing to protect its territory. Stranded and under constant attack, Paul and his family must somehow find a way to get his family back to shore alive before it strikes again in this epic battle between humans and nature.",
			ImagePath: 	 "images/movie_images/890771",
		},
		
	
		{
			Title: 		 "Nimona",
			Slug: 		 "a-new-hero-takes-shapei",
			Year: 		 "2023",
			Description: 	 "A knight framed for a tragic crime teams with a scrappy, shape-shifting teen to prove his innocence.",
			ImagePath: 	 "images/movie_images/961323",
		},
		
	
		{
			Title: 		 "Sound of Freedom",
			Slug: 		 "fight-for-the-lighti-silence-the-darknessi",
			Year: 		 "2023",
			Description: 	 "The story of Tim Ballard, a former US government agent, who quits his job in order to devote his life to rescuing children from global sex traffickers.",
			ImagePath: 	 "images/movie_images/678512",
		},
		
	
		{
			Title: 		 "The Out-Laws",
			Slug: 		 "meeting-the-in-laws-has-never-been-this-dangerousi",
			Year: 		 "2023",
			Description: 	 "A straight-laced bank manager is about to marry the love of his life. When his bank is held up by infamous Ghost Bandits during his wedding week, he believes his future in-laws who just arrived in town, are the infamous Out-Laws.",
			ImagePath: 	 "images/movie_images/921636",
		},
		
	
		{
			Title: 		 "Boston Strangler",
			Slug: 		 "inspired-by-a-true-story",
			Year: 		 "2023",
			Description: 	 "Reporters Loretta McLaughlin and Jean Cole bravely pursue the story of the Boston Strangler at great personal risk, putting their own lives on the line in their quest to uncover the truth.",
			ImagePath: 	 "images/movie_images/881164",
		},
		
	
		{
			Title: 		 "We Have a Ghost",
			Slug: 		 "set-your-spirit-freei",
			Year: 		 "2023",
			Description: 	 "After Kevin finds a ghost named Ernest haunting his new home, he becomes an overnight social media sensation. But when Kevin and Ernest go rogue to investigate the mystery of the latter's past, they become targets of the CIA.",
			ImagePath: 	 "images/movie_images/852096",
		},
		
	
		{
			Title: 		 "Infinity Pool",
			Slug: 		 "find-out-what-kind-of-a-creature-you-really-arei",
			Year: 		 "2023",
			Description: 	 "While staying at an isolated island resort, James and Em are enjoying a perfect vacation of pristine beaches, exceptional staff, and soaking up the sun. But guided by the seductive and mysterious Gabi, they venture outside the resort grounds and find themselves in a culture filled with violence, hedonism, and untold horror. A tragic accident leaves them facing a zero tolerance policy for crime: either you'll be executed, or, if you’re rich enough to afford it, you can watch yourself die instead.",
			ImagePath: 	 "images/movie_images/667216",
		},
		
	
		{
			Title: 		 "The Boogeyman",
			Slug: 		 "it-s-not-reali-it-s-not-reali-it-s-not-reali",
			Year: 		 "2023",
			Description: 	 "Still reeling from the tragic death of their mother, a teenage girl and her younger sister find themselves plagued by a sadistic presence in their house and struggle to get their grieving father to pay attention before it’s too late.",
			ImagePath: 	 "images/movie_images/532408",
		},
		
	
		{
			Title: 		 "Flamin' Hot",
			Slug: 		 "the-flavor-you-knowi-the-story-you-don-ti",
			Year: 		 "2023",
			Description: 	 "The inspiring true story of Richard Montañez, the Frito Lay janitor who channeled his Mexican American heritage and upbringing to turn the iconic Flamin’ Hot Cheetos into a snack that disrupted the food industry and became a global pop culture phenomenon.",
			ImagePath: 	 "images/movie_images/626332",
		},
		
	
		{
			Title: 		 "Cobweb",
			Slug: 		 "sooner-or-lateri-family-secrets-creep-outi",
			Year: 		 "2023",
			Description: 	 "Eight year old Peter is plagued by a mysterious, constant tapping from inside his bedroom wall—one that his parents insist is all in his imagination. As Peter's fear intensifies, he believes that his parents could be hiding a terrible, dangerous secret and questions their trustworthiness.",
			ImagePath: 	 "images/movie_images/709631",
		},
		
	
		{
			Title: 		 "Strays",
			Slug: 		 "go-fetch-yourselfi",
			Year: 		 "2023",
			Description: 	 "When Reggie is abandoned on the mean city streets by his lowlife owner, Doug, Reggie is certain that his beloved owner would never leave him on purpose. But once Reggie falls in with Bug, a fast-talking, foul-mouthed stray who loves his freedom and believes that owners are for suckers, Reggie finally realizes he was in a toxic relationship and begins to see Doug for the heartless sleazeball that he is.",
			ImagePath: 	 "images/movie_images/912908",
		},
		
	
		{
			Title: 		 "The Nun II",
			Slug: 		 "confess-your-sinsi",
			Year: 		 "2023",
			Description: 	 "In 1956 France, a priest is violently murdered, and Sister Irene begins to investigate. She once again comes face-to-face with a powerful evil.",
			ImagePath: 	 "images/movie_images/968051",
		},
		
	
		{
			Title: 		 "Magic Mike's Last Dance",
			Slug: 		 "the-final-teasei",
			Year: 		 "2023",
			Description: 	 "Mike Lane takes to the stage again after a lengthy hiatus, following a business deal that went bust, leaving him broke and taking bartender gigs in Florida. For what he hopes will be one last hurrah, Mike heads to London with a wealthy socialite who lures him with an offer he can’t refuse… and an agenda all her own. With everything on the line, once Mike discovers what she truly has in mind, will he—and the roster of hot new dancers he’ll have to whip into shape—be able to pull it off?",
			ImagePath: 	 "images/movie_images/906221",
		},
		
	
		{
			Title: 		 "They Cloned Tyrone",
			Slug: 		 "beware-cheap-imitationsi",
			Year: 		 "2023",
			Description: 	 "A series of eerie events thrusts an unlikely trio onto the trail of a nefarious government conspiracy lurking directly beneath their neighborhood.",
			ImagePath: 	 "images/movie_images/736769",
		},
		
	
		{
			Title: 		 "Sharper",
			Slug: 		 "read-between-the-liesi",
			Year: 		 "2023",
			Description: 	 "A small, wealthy family in New York City gets progressively torn apart by secrets, lies, and the theft that orchestrates all of it.",
			ImagePath: 	 "images/movie_images/717980",
		},
		
	
		{
			Title: 		 "No One Will Save You",
			Slug: 		 "a-home-invasion-no-one-saw-comingi",
			Year: 		 "2023",
			Description: 	 "An exiled anxiety-ridden homebody must battle an alien who's found its way into her home.",
			ImagePath: 	 "images/movie_images/820609",
		},
		
	
		{
			Title: 		 "Die Hart",
			Slug: 		 "live-to-change",
			Year: 		 "2023",
			Description: 	 "Kevin Hart - playing a version of himself - is on a death-defying quest to become an action star. And with a little help from John Travolta, Nathalie Emmanuel, and Josh Hartnett - he just might pull it off.",
			ImagePath: 	 "images/movie_images/1077280",
		},
		
	
		{
			Title: 		 "Chupa",
			Slug: 		 "discover-the-legendi",
			Year: 		 "2023",
			Description: 	 "While visiting family in Mexico, a lonely boy befriends a mythical creature hiding on his grandfather's ranch and embarks on the adventure of a lifetime.",
			ImagePath: 	 "images/movie_images/736790",
		},
		
	
		{
			Title: 		 "The Devil Conspiracy",
			Slug: 		 "the-war-of-angels-has-come-to-earthi",
			Year: 		 "2023",
			Description: 	 "A powerful biotech company has breakthrough technology allowing them to clone history’s most influential people with just a few fragments of DNA. Behind this company is a cabal of Satanists that steals the shroud of Christ, putting them in possession of Jesus’ DNA. The clone will serve as the ultimate offering to the devil. The Archangel Michael comes to earth and will stop at nothing to end the devil’s conspiracy.",
			ImagePath: 	 "images/movie_images/296271",
		},
		
	
		{
			Title: 		 "Retribution",
			Slug: 		 "all-roads-lead-to-the-truthi",
			Year: 		 "2023",
			Description: 	 "When a mysterious caller puts a bomb under his car seat, Matt Turner begins a high-speed chase across the city to complete a specific series of tasks. With his kids trapped in the back seat and a bomb that will explode if they get out of the car, a normal commute becomes a twisted game of life or death as Matt follows the stranger's increasingly dangerous instructions in a race against time to save his family.",
			ImagePath: 	 "images/movie_images/762430",
		},
		
	
		{
			Title: 		 "The Equalizer 3",
			Slug: 		 "justice-knows-no-bordersi",
			Year: 		 "2023",
			Description: 	 "Robert McCall finds himself at home in Southern Italy but he discovers his friends are under the control of local crime bosses. As events turn deadly, McCall knows what he has to do: become his friends' protector by taking on the mafia.",
			ImagePath: 	 "images/movie_images/926393",
		},
		
	
		{
			Title: 		 "Marlowe",
			Slug: 		 "los-angeles-1939i-outside-the-spotlight-lies-a-city-of-secretsi",
			Year: 		 "2023",
			Description: 	 "Private detective Philip Marlowe becomes embroiled in an investigation involving a wealthy Californian family after a beautiful blonde hires him to track down her former lover.",
			ImagePath: 	 "images/movie_images/844417",
		},
		
	
		{
			Title: 		 "Haunted Mansion",
			Slug: 		 "home-is-where-the-haunt-isi",
			Year: 		 "2023",
			Description: 	 "A woman and her son enlist a motley crew of so-called spiritual experts to help rid their home of supernatural squatters.",
			ImagePath: 	 "images/movie_images/616747",
		},
		
	
		{
			Title: 		 "Mighty Morphin Power Rangers: Once & Always",
			Slug: 		 "a-mighty-morphin-reunion-30-years-in-the-makingi",
			Year: 		 "2023",
			Description: 	 "After tragedy strikes, an unlikely young hero takes her rightful place among the Power Rangers to face off against the team's oldest archnemesis.",
			ImagePath: 	 "images/movie_images/1068141",
		},
		
	
		{
			Title: 		 "Crater",
			Slug: 		 "miles-from-earthi-their-adventure-beginsi",
			Year: 		 "2023",
			Description: 	 "After the death of his father, a boy growing up on a lunar mining colony takes a trip to explore a legendary crater, along with his four best friends, prior to being permanently relocated to another planet.",
			ImagePath: 	 "images/movie_images/620705",
		},
		
	
		{
			Title: 		 "Love at First Sight",
			Slug: 		 "100i000-daily-flights-around-the-worldi-6-million-travelersi-one-connectioni",
			Year: 		 "2023",
			Description: 	 "Hadley and Oliver begin falling in love on a flight from New York to London, but when they lose each other at customs, can they defy all odds to reunite?",
			ImagePath: 	 "images/movie_images/353577",
		},
		
	
		{
			Title: 		 "Champions",
			Slug: 		 "every-dream-team-starts-somewherei",
			Year: 		 "2023",
			Description: 	 "A stubborn and hotheaded minor league basketball coach is forced to train a Special Olympics team when he is sentenced to community service.",
			ImagePath: 	 "images/movie_images/933419",
		},
		
	
		{
			Title: 		 "White Men Can't Jump",
			Slug: 		 "play-hardi-hustle-harderi",
			Year: 		 "2023",
			Description: 	 "Seemingly opposite street hoopers, Jeremy, an injury prone former star, and Kamal, a has-been prodigy, team up to take one final shot at living out their dreams.",
			ImagePath: 	 "images/movie_images/920125",
		},
		
	
		{
			Title: 		 "Inside",
			Slug: 		 "a-solitary-exhibitioni",
			Year: 		 "2023",
			Description: 	 "An art thief becomes trapped in a New York penthouse after his heist goes awry. Imprisoned with nothing but priceless works of art, he must use all his cunning and invention to survive.",
			ImagePath: 	 "images/movie_images/958196",
		},
		
	
		{
			Title: 		 "Pamela, A Love Story",
			Slug: 		 "she-defined-a-decadei-now-she-will-define-herselfi",
			Year: 		 "2023",
			Description: 	 "In her own words, through personal video and diaries, Pamela Anderson shares the story of her rise to fame, rocky romances and infamous sex tape scandal.",
			ImagePath: 	 "images/movie_images/1061671",
		},
		
	
		{
			Title: 		 "Love Again",
			Slug: 		 "destiny-has-a-plani",
			Year: 		 "2023",
			Description: 	 "Mira Ray, dealing with the loss of her fiancé, sends a series of romantic texts to his old cell phone number… not realizing the number was reassigned to Rob Burns' new work phone. A journalist, Rob is captivated by the honesty in the beautifully confessional texts. When he’s assigned to write a profile of megastar Céline Dion, he enlists her help in figuring out how to meet Mira in person and win her heart.",
			ImagePath: 	 "images/movie_images/758336",
		},
		
	
		{
			Title: 		 "A Good Person",
			Slug: 		 "sometimes-we-find-hope-where-we-least-expect-iti",
			Year: 		 "2023",
			Description: 	 "Allison's life falls apart following her involvement in a fatal accident. The unlikely relationship she forms with her would-be father-in-law helps her live a life worth living.",
			ImagePath: 	 "images/movie_images/800787",
		},
		
	
		{
			Title: 		 "BlackBerry",
			Slug: 		 "work-hardi-fail-hardi",
			Year: 		 "2023",
			Description: 	 "Two mismatched entrepreneurs – egghead innovator Mike Lazaridis and cut-throat businessman Jim Balsillie – joined forces in an endeavour that was to become a worldwide hit in little more than a decade. The story of the meteoric rise and catastrophic demise of the world's first smartphone.",
			ImagePath: 	 "images/movie_images/1016084",
		},
		
	
		{
			Title: 		 "Warhorse One",
			Slug: 		 "worth-living-fori-worth-dying-fori",
			Year: 		 "2023",
			Description: 	 "A gunned down Navy SEAL Master Chief must guide a child to safety through a gauntlet of hostile Taliban insurgents and survive the brutal Afghanistan wilderness.",
			ImagePath: 	 "images/movie_images/1076487",
		},
		
	
		{
			Title: 		 "A Million Miles Away",
			Slug: 		 "it-takes-more-than-a-dream-to-reach-the-starsi",
			Year: 		 "2023",
			Description: 	 "The life of engineer and former NASA astronaut José M. Hernández, the first migrant farmworker to go to space.",
			ImagePath: 	 "images/movie_images/1002185",
		},
		
	
		{
			Title: 		 "Reality",
			Slug: 		 "the-truth-cannot-be-redacted",
			Year: 		 "2023",
			Description: 	 "Augusta, Georgia, United States, June 3, 2017. After running some errands, Reality Winner returns home, where she is approached by two men.",
			ImagePath: 	 "images/movie_images/985617",
		},
		
	
		{
			Title: 		 "The Monkey King",
			Slug: 		 "the-legend-has-arrived",
			Year: 		 "2023",
			Description: 	 "A stick-wielding monkey teams with a young girl on an epic quest for immortality, battling demons, dragons, gods — and his own ego — along the way.",
			ImagePath: 	 "images/movie_images/832502",
		},
		
	
		{
			Title: 		 "Joy Ride",
			Slug: 		 "four-friendsi-one-tripi-no-lucki",
			Year: 		 "2023",
			Description: 	 "When Audrey's business trip to Asia goes sideways, she enlists the aid of Lolo, her irreverent, childhood best friend who also happens to be a hot mess; Kat, her college friend turned Chinese soap star; and Deadeye, Lolo's eccentric cousin. Their no-holds-barred, epic experience becomes a journey of bonding, friendship, belonging, and wild debauchery that reveals the universal truth of what it means to know and love who you are.",
			ImagePath: 	 "images/movie_images/864168",
		},
		
	
		{
			Title: 		 "Justice League: Warworld",
			Slug: 		 "three-heroesi-three-worldsi-one-salvationi",
			Year: 		 "2023",
			Description: 	 "Until now, the Justice League has been a loose association of superpowered individuals. But when they are swept away to Warworld, a place of unending brutal gladiatorial combat, Batman, Superman, Wonder Woman and the others must somehow unite to form an unbeatable resistance able to lead an entire planet to freedom.",
			ImagePath: 	 "images/movie_images/1003581",
		},
		
	
		{
			Title: 		 "Skinamarink",
			Slug: 		 "in-this-houseiii",
			Year: 		 "2023",
			Description: 	 "Two children wake up in the middle of the night to find their father is missing, and all the windows and doors in their home have vanished.",
			ImagePath: 	 "images/movie_images/994143",
		},
		
	
		{
			Title: 		 "Clock",
			Slug: 		 "it99s-counting-down-for-a-reasoni",
			Year: 		 "2023",
			Description: 	 "On the eve of her 38th birthday, a woman desperately attempts to fix her broken biological clock.",
			ImagePath: 	 "images/movie_images/1085103",
		},
		
	
		{
			Title: 		 "Vacation Friends 2",
			Slug: 		 "time-for-another-roundi",
			Year: 		 "2023",
			Description: 	 "Newly married couple Marcus and Emily invite their uninhibited besties Ron and Kyla to join them for a vacation when Marcus lands an all-expenses-paid trip to a Caribbean resort. When Kyla’s incarcerated father Reese is released and shows up at the resort unannounced, things get out of control, upending Marcus’ best laid plans and turning the vacation friends’ perfect trip into total chaos.",
			ImagePath: 	 "images/movie_images/869641",
		},
		
	
		{
			Title: 		 "Dog Gone",
			Slug: 		 "to-find-what-s-losti-a-father-and-son-will-do-the-impossiblei",
			Year: 		 "2023",
			Description: 	 "When his beloved dog goes missing, a young man embarks on an incredible search with his parents to find him and give him life-saving medication.",
			ImagePath: 	 "images/movie_images/858408",
		},
		
	
		{
			Title: 		 "Are You There God? It's Me, Margaret.",
			Slug: 		 "discovering-who-you-are-is-a-journey-that-lasts-a-lifetimei",
			Year: 		 "2023",
			Description: 	 "When her family moves from New York City to New Jersey, an 11-year-old girl navigates new friends, feelings, and the beginning of adolescence.",
			ImagePath: 	 "images/movie_images/555285",
		},
		
	
		{
			Title: 		 "Big George Foreman",
			Slug: 		 "when-life-gives-you-a-second-chancei-don-t-do-it-for-yourselfi-do-it-for-everyone-you-lovei",
			Year: 		 "2023",
			Description: 	 "Fueled by an impoverished childhood, George Foreman channeled his anger into becoming an Olympic Gold medalist and World Heavyweight Champion, followed by a near-death experience that took him from the boxing ring to the pulpit. But when he sees his community struggling spiritually and financially, Foreman returns to the ring and makes history by reclaiming his title, becoming the oldest and most improbable World Heavyweight Boxing Champion ever.",
			ImagePath: 	 "images/movie_images/878361",
		},
		
	
		{
			Title: 		 "Mafia Mamma",
			Slug: 		 "from-suburban-mom-to-mafia-doni",
			Year: 		 "2023",
			Description: 	 "A suburban American woman inherits her grandfather’s Mafia empire and, guided by the Firm’s trusted consigliere, defies everyone’s expectations, including her own, as the new head of the family business.",
			ImagePath: 	 "images/movie_images/809787",
		},
	}

	if err := db.Create(&movies).Error; err != nil {
		log.Fatalf("Failed to seed genres table: %v", err)
	}
	log.Println("Genres table seeded successfully")
}

func SeedMovieGenresActors(db *gorm.DB) {
    // Fetch all Genre IDs
    var genreIDs []uint
    if err := db.Model(&domain.Genre{}).Pluck("id", &genreIDs).Error; err != nil {
        log.Fatalf("Failed to fetch genres: %v", err)
    }

    // Fetch all Actor IDs
    var actorIDs []uint
    if err := db.Model(&domain.Actor{}).Pluck("id", &actorIDs).Error; err != nil {
        log.Fatalf("Failed to fetch actors: %v", err)
    }

    log.Printf("Fetched %d Genre IDs: %v", len(genreIDs), genreIDs)
    log.Printf("Fetched %d Actor IDs: %v", len(actorIDs), actorIDs)

    // Seed random data into the Movies table
	var movies []domain.Movie
	if err := db.Where("genre_id IS NULL OR actor_id IS NULL").Find(&movies).Error; err != nil {
		log.Fatalf("Failed to fetch movies: %v", err)
	}

    // Update each movie with random genre_id and actor_id
    for i := range movies {
        movies[i].GenreID = &genreIDs[rand.Intn(len(genreIDs))]
        movies[i].ActorID = &actorIDs[rand.Intn(len(actorIDs))]

        // Save the updated movie
        if err := db.Save(&movies[i]).Error; err != nil {
            log.Printf("Failed to update movie ID %d: %v", movies[i].ID, err)
        } else {
            log.Printf("Successfully updated movie ID %d with GenreID %d and ActorID %d", movies[i].ID, movies[i].GenreID, movies[i].ActorID)
        }
    }

    log.Println("All movies updated successfully!")
}
