# knavigator
Generate a Knave character

### How to ...

- Execute linting and unit tests:
```
make test
```

- Run locally as command:
```
make run
```

- Run locally as web app:
```
make serve
```

( Visit `http://localhost:3000/character` in the browser )

- Build locally:
```
make build
```

### What is Knave?

[Available at DriveThruRPG](https://www.drivethrurpg.com/product/250888/Knave?affiliate_id=379088), Knave is a rules-lite toolkit for running old school fantasy RPGs without classes.

### Sample Web Output

![screenshot](/assets/screenshot.png?raw=true "Screenshot")

### Sample Command Output

```
$ make run
go run cmd/app/main.go

Chloe Moon
Strength: 13 / +3
Dexterity: 11 / +1
Constitution: 12 / +2
Intelligence: 11 / +1
Wisdom: 11 / +1
Charisma: 11 / +1
HP: 8 / +4
Defense: 13 / +3
Weapon: Cudgel (d6 damage, 1 slots, 1 hands, 3 quality)
Armor: Brigandine (Defense 13, 2 Slots, Quality 4)
Dungeoneering Gear: [Candles, 5 (1 Slots)]
General Gear: [Rations (1 Slots) Rations (1 Slots) Bear trap (1 Slots) Lens (1 Slots)]
Slots: 7 / 12
Optional Starting Money: 220 copper pieces
Traits: Has a LANKY physique with a WIDE face, PIERCED skin, and GREASED hair. Wears UNDERSIZED clothing and speaks in a DRAWLING voice. Is HONORABLE, yet IRASCIBLE. Aligned towards NEUTRALITY. Was once a BUTCHER, and has had the misfortune of being CONDEMNED.
```

## Development

- TODO: Unit tests !!!
- TODO: Turning resulting data into nice-looking presentation

### Thanks

- Thanks to https://lawbreaker.herokuapp.com/ for inspiring me and posting useful code.
- Thanks to [u/OrkishBlade](https://www.reddit.com/r/DnDBehindTheScreen/comments/50pcg1/a_post_about_names_names_for_speakers_of_the/) for the random names.
