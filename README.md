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

Henro Pottaker
Strength: 11 / +1
Dexterity: 14 / +4
Constitution: 13 / +3
Intelligence: 11 / +1
Wisdom: 11 / +1
Charisma: 11 / +1
HP: 7 / Healing Rate: 8
Defense: 11 / +1
Armor: none
Dungeoneering Gear: []
General Gear: []
Traits: Has a TINY physique with a DELICATE face, BURN SCAR skin, and DREADLOCKS hair. Wears PATCHED clothing and speaks in a RAPID-FIRE voice. Is CURIOUS and believes in CHAOS, yet GREEDY. Was once a PICKPOCKET, but has since come to be HAUNTED.
```

## Development

- TODO: Unit tests !!!
- TODO: Generate armor
- TODO: Generate gear
- TODO: Turning resulting data into nice-looking presentation

### Thanks

- Thanks to https://lawbreaker.herokuapp.com/ for inspiring me and posting useful code.
- Thanks to [u/OrkishBlade](https://www.reddit.com/r/DnDBehindTheScreen/comments/50pcg1/a_post_about_names_names_for_speakers_of_the/) for the random names.
