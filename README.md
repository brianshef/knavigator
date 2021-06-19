# knavigator
Generate a Knave character

### How to ...

- Execute linting and unit tests:
```
make test
```

- Run locally:
```
make run
```

- Build locally:
```
make build
```

### What is Knave?

[Available at DriveThruRPG](https://www.drivethrurpg.com/product/250888/Knave?affiliate_id=379088), Knave is a rules-lite toolkit for running old school fantasy RPGs without classes.

### Sample Output

```
$ make run
go run cmd/app/main.go

Maya Oakes
Strength: 12 / +2
Dexterity: 11 / +1
Constitution: 11 / +1
Intelligence: 11 / +1
Wisdom: 11 / +1
Charisma: 11 / +1
HP: 6 / +7
Defense: 14 / +4
Weapon: Dagger (d6 damage, 1 slots, 1 hands, 3 quality)
Armor: Brigandine with helmet (Defense 14, 3 Slots, Quality 4)
Dungeoneering Gear: []
General Gear: []
Slots: 7 / 11
Traits: Has a GAUNT physique with a PATRICIAN face, PIERCED skin, and BALD hair. Wears OVERSIZED clothing and speaks in a RAPID-FIRE voice. Is MERCIFUL, yet CRUEL. Aligned towards LAW. Was once a GAMBLER, and has had the misfortune of being EXILED.
```

## Development

- TODO: Unit tests !!!
- TODO: Generate gear
- TODO: Turning resulting data into nice-looking presentation

### Thanks

- Thanks to https://lawbreaker.herokuapp.com/ for inspiring me and posting useful code.
- Thanks to [u/OrkishBlade](https://www.reddit.com/r/DnDBehindTheScreen/comments/50pcg1/a_post_about_names_names_for_speakers_of_the/) for the random names.
