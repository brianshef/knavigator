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

Aliya Redwine
Strength: 11 / +1
Dexterity: 12 / +2
Constitution: 12 / +2
Intelligence: 14 / +4
Wisdom: 14 / +4
Charisma: 13 / +3
HP: 7 / +8
Defense: 14 / +4
Armor: Brigandine with shield (Defense 14, 3 Slots, Quality 4)
Dungeoneering Gear: []
General Gear: []
Slots: 9 / 12
Traits: Has a STATUESQUE physique with a SQUARE face, POCKMARKED skin, and CROPPED hair. Wears FLAMBOYANT clothing and speaks in a GRAVELLY voice. Is DISCIPLINED, yet WASTEFUL. Aligned towards LAW. Was once a CHARLATAN, and has had the misfortune of being HAUNTED.
```

## Development

- TODO: Unit tests !!!
- TODO: Generate gear
- TODO: Turning resulting data into nice-looking presentation

### Thanks

- Thanks to https://lawbreaker.herokuapp.com/ for inspiring me and posting useful code.
- Thanks to [u/OrkishBlade](https://www.reddit.com/r/DnDBehindTheScreen/comments/50pcg1/a_post_about_names_names_for_speakers_of_the/) for the random names.
