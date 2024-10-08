# KancoLaTeX

A command line tool generate LaTex template from Kancolle Deck Builder.

## Requirement

- Python 3.10 or greater

## How to setup

0. Launch a terminal

<details open>
    <summary>Download from pip</summary>

1. Create a python venv

```bash
python3 -m venv .venv
```

2. Enable venv
   - on Unix system
   ```sh
   source .venv/bin/activate
   ```

   - on Windows
   ```
   .\.venv\Scripts\activate
   ```

   When the successfully launch, there should be a `(.venv)` on your terminal.

   > **IT IS STRONGLY ENCOURAGE TO CREATE A VIRTUAL ENVIRONMENT FOR ANY PYTHON
   > PROJECT WITH 3RD PARTY DEPENDENCE!**

   > use `deactivate` to quit the venv.
   > ```sh
   > deactivate
   > ```

3. Download the package in pip

   ```sh
   pip install kancolatex
   ```
4. Test is install successful After the setup. when you type `kancolatex` into
   your terminal, it shall display the help into.

   ```sh
   kancolatex
   ```

5. If it success, next you shall init a database for game data by enter:

   ```sh
   kancolatex --reset
   ```
</details>

<details>
    <summary>Download from source</summary>

1. Clone this repo

- If you don't know git, enter following command into your terminal

  ```sh
  git clone https://github.com/kafe523/kancolatex.git
  ```

  After cloning, please cd inside the folder

  ```sh
  cd kancolatex/
  ```

1. Create a python venv

   ```bash
   python3 -m venv .venv
   ```

2. Enable venv
   - on Unix system

   ```sh
   source .venv/bin/activate
   ```

   - on Windows

   ```
   .\.venv\Scripts\activate
   ```

   When the successfully launch, there should be a `(.venv)` on your terminal.

   > **IT IS STRONGLY ENCOURAGE TO CREATE A VIRTUAL ENVIRONMENT FOR ANY PYTHON
   > PROJECT WITH 3RD PARTY DEPENDENCE!**

   > use `deactivate` to quit the venv.
   > ```sh
   > deactivate
   > ```

3. Install dependence After enable the venv, type following command to install
   dependence.

   ```sh
   pip install -e .
   ```


4. Test is install successful After the setup. when you type `kancolatex` into
   your terminal, it shall display the help into.

   ```sh
   kancolatex
   ```

5. If it success, next you shall init a database for game data by enter:

   ```sh
   kancolatex --reset
   ```
</details>

## How to update

<details open>
    <summary>Download from pip</summary>

`pip install kancolatex --upgrade`

</details>

<details>
    <summary>Download from source</summary>
    
1. Deactivate your python venv is it is enabled
2. Delete your local `Kancolatex` folder
3. [How to setup](#how-to-setup)

</details>

## How to use

The program have 3 working mode.

1. [Default Mode](#default-mode)
2. [Export Mode](#export-mode)
3. [Translate Mode](#translate-mode)

You can change the mode with `-m / --mode` flag.

Example:

- `kancolatex` (default is Default Mode)
- `kancolatex -m default` (Which will configure to Default Mode)
- `kancolatex --mode export` (Which will configure to Export Mode)
- `kancolatex --mode translate` (Which will configure to Translate Mode)

### Default Mode

This is the default mode of the kancolatex. Which is mode working before version 0.0.5 .

Which will parse then fill with the corresponded value from deck builder into latex template.

You must prepare:

- An deck json
- A latex template

#### Preprocess Mode

If you want to access deck json data from latex template. The latex template
should met such criteria:

1. The program only evaluate the macro inside `KancoLaTeX:preprocess` scope. You
   can start the preprocess scope by written:
   ```latex
   % KancoLaTeX:preprocess:begin
   ```

   closing preprocess scope by written:

   ```latex
   % KancoLaTeX:preprocess:end
   ```
   > **These must occupying single line without any other character in the same
   > line. And the `%` must be the first letter of the line.**

2. Inside the scope of, you can write [Predefined macro](#predefined-macro) to
   access the data you want.

#### Define Mode

(To Be Finish)

- define json schema

```ts
type DefineConfig = {
    overwrite?: boolean;
    name: string;
    template: string;
    param?: string[];
};
```

- Macro mode `<...>`
- Access mode `<<...>>`

- Example

```latex
% KancoLaTeX:define:begin

% {"name": "\\foo", "template": "\\bar{{{0}}}", "param": ["abc"]}
% {"name": "\\ham", "template": "\\egg{{{0}}}{{{1}}}", "param": ["<SHIP_A_NAME_EN>", "<<Ship.A.level>>"]}

% KancoLaTeX:define:end
```

#### Output

Default the the result will output to `stdout` and log to `stderr`. If you want
to store the result somewhere else can use `-o` flag or shell pipeline.

#### Example 1

```sh
kancolatex -t ./example_1_template.tex -n ./example_1_fleet.json
```

- `example_1_fleet.json`

```json
{
    "version": 4,
    "hqlv": 120,
    "f1": {
        "name": "第1艦隊",
        "t": 0,
        "s1": {
            "id": 610,
            "lv": 99,
            "exa": true,
            "items": {
                "i1": { "id": 343, "rf": 2, "mas": 4 },
                "i2": { "id": 100, "rf": 0, "mas": 7 },
                "i3": { "id": 157, "rf": 10, "mas": 7 },
                "i4": { "id": 339, "rf": 0, "mas": 7 },
                "i5": { "id": 273, "rf": 2, "mas": 6 },
                "ix": { "id": 478, "rf": 9 }
            },
            "hp": 84,
            "asw": 13,
            "luck": 18
        }
    },
    "f2": {},
    "f3": {},
    "f4": {},
    "a1": {
        "name": "第1基地航空隊",
        "items": {
            "i1": { "id": 263, "rf": 0, "mas": 7 },
            "i2": { "id": 264, "rf": 2, "mas": 6 },
            "i3": { "id": 224, "rf": 0, "mas": 5 },
            "i4": { "id": 404, "rf": 0, "mas": 3 }
        },
        "mode": 1
    },
    "a2": {
        "name": "第2基地航空隊",
        "items": {
            "i1": { "id": 479, "rf": 0, "mas": 4 },
            "i2": { "id": 264, "rf": 0, "mas": 7 },
            "i3": { "id": 224, "rf": 0, "mas": 5 },
            "i4": { "id": 493, "rf": 4, "mas": 4 }
        },
        "mode": 1
    },
    "a3": {
        "name": "第3基地航空隊",
        "items": {
            "i1": { "id": 311, "rf": 0, "mas": 7 },
            "i2": { "id": 265, "rf": 0, "mas": 7 },
            "i3": { "id": 454, "rf": 0, "mas": 5 },
            "i4": { "id": 454, "rf": 0, "mas": 5 }
        },
        "mode": 1
    },
    "s": {
        "a": 0,
        "i": 0,
        "c": [{ "c": 0, "pf": 1, "ef": 1, "f1": { "s": [] } }]
    }
}
```

- `example_1_template.tex`

```tex
% KancoLaTeX:preprocess:begin
\begin{document}
% wrote your latex template here
% kancolatex will and only replace the symbol inside the scope.

\begin{tblr}{colspec={Q[c,m]Q[c,m]},vlines,hlines}
% ship name japanese & ship name english & ship level
\shipAnameJp & \shipAnameEn & \shipAlevel & \\
\end{tblr}

\begin{tblr}{colspec={Q[c,m]Q[c,m]},vlines,hlines}
% equipment name japaese & equipment name english & equipment remodel & equipment aircraft level
\shipAequipmentAnameJp & \shipAequipmentAnameEn & \shipAequipmentAremodel & \shipAequipmentAlevelAlt \\
\shipAequipmentBnameJp & \shipAequipmentBnameEn & \shipAequipmentBremodel & \shipAequipmentBlevelAlt \\
\shipAequipmentCnameJp & \shipAequipmentCnameEn & \shipAequipmentCremodel & \shipAequipmentClevelAlt \\
\shipAequipmentDnameJp & \shipAequipmentDnameEn & \shipAequipmentDremodel & \shipAequipmentDlevelAlt \\
\shipAequipmentEnameJp & \shipAequipmentEnameEn & \shipAequipmentEremodel & \shipAequipmentElevelAlt \\
\shipAequipmentXnameJp & \shipAequipmentXnameEn & \shipAequipmentXremodel & \shipAequipmentXlevelAlt \\
\end{tblr}

\end{document}
% KancoLaTeX:preprocess:end
```

- output

```tex
% KancoLaTeX:preprocess:begin
\begin{document}
% wrote your latex template here
% kancolatex will and only replace the symbol inside the scope.

\begin{tblr}{colspec={Q[c,m]Q[c,m]},vlines,hlines}
% ship name japanese & ship name english & ship level
加賀改二戊 & Kaga Kai Ni E & 99 & \\
\end{tblr}

\begin{tblr}{colspec={Q[c,m]Q[c,m]},vlines,hlines}
% equipment name japaese & equipment name english & equipment remodel & equipment aircraft level
流星改(一航戦/熟練) & Ryuusei Kai (CarDiv 1/Skilled) & 2 & 4 \\
彗星(江草隊) & Suisei (Egusa Squadron) & 0 & 7 \\
零式艦戦53型(岩本隊) & Type 0 Fighter Model 53 (Iwamoto Squadron) & 10 & 7 \\
烈風改二戊型(一航戦/熟練) & Reppuu Kai 2 Model E (CarDiv 1/Skilled) & 0 & 7 \\
彩雲(偵四) & Saiun (4th Recon Squad) & 2 & 6 \\
熟練甲板要員+航空整備員 & Skilled Deck Personnel + Aviation Maintenance Hands & 9 & 0 \\
\end{tblr}

\end{document}
% KancoLaTeX:preprocess:end
```

#### Predefined Macro

- Airbase

| Mnemonic | Airbase     |
| -------- | ----------- |
| A        | airbase1    |
| B        | airbase2    |
| C        | airbase1    |
| U        | All airbase |

> `P = airbase mnemonic`

| Latex                            | Macro                               | Access                            | Usage                 | Example                        | Note                                                                  |
| -------------------------------- | ----------------------------------- | --------------------------------- | --------------------- | ------------------------------ | --------------------------------------------------------------------- |
| `\airbase{P}fullAirPower`        | `AIRBASE_{P}_FULL_AIRPOWER`         | `Airbase.{P}.fullAirPower`        | Sortie Airpower       | `\airbaseAfullAirPower`        | Will be empty if airbase in defense mode                              |
| `\airbase{P}defenseAirPower`     | `AIRBASE_{P}_DEFENSE_AIRPOWER`      | `Airbase.{P}.defenseAirPower`     | Defense Airpower      | `\airbaseBdefenseAirPower`     |                                                                       |
| `\airbase{P}highDefenseAirPower` | `AIRBASE_{P}_HIGH_DEFENSE_AIRPOWER` | `Airbase.{P}.highDefenseAirPower` | High Defense Airpower | `\airbaseChighDefenseAirPower` |                                                                       |
| `\airbase{P}mode`                | `AIRBASE_{P}_MODE`                  | `Airbase.{P}.mode`                | Airbase mode          | `\airbaseAmode`                | 0 = Standby, 1 = Mission, 2 = Defense; airbase U access will be empty |
| `\airbase{P}radius`              | `AIRBASE_{P}_RADIUS`                | `Airbase.{P}.radius`              | Airbase radius        | `\airbaseBradius`              | airbase U access will be empty                                        |

- Airbase Equipment

| Mnemonic | Airbase              |
| -------- | -------------------- |
| A        | airbase equipment  1 |
| B        | airbase equipment  2 |
| C        | airbase equipment 3  |
| D        | airbase equipment 4  |

> `P = airbase mnemonic` `T = equipment mnemonic`

| Latex                             | Macro                                 | Access                                      | Usage                      | Example                        | Note                             |
| --------------------------------- | ------------------------------------- | ------------------------------------------- | -------------------------- | ------------------------------ | -------------------------------- |
| `\airbase{P}equipment{T}nameJp`   | `AIRBASE_{P}_EQUIPMENT_{T}_NAME_JP`   | `AIRBASE.{P}.equipment.{T}.data.name`       | Equipment Japanese name    | `\airbaseAequipmentAnameJp`    |                                  |
| `\airbase{P}equipment{T}nameEn`   | `AIRBASE_{P}_EQUIPMENT_{T}_NAME_EN`   | ``                                          | Equipment English name     | `\airbaseBequipmentXnameEn`    | No Access Method                 |
| `\airbase{P}equipment{T}remodel`  | `AIRBASE_{P}_EQUIPMENT_{T}_REMODEL`   | `AIRBASE.{P}.equipment.{T}remodel`          | Equipment level of remodel | `\airbaseCequipmentCremodel`   |                                  |
| `\airbase{P}equipment{T}levelAlt` | `AIRBASE_{P}_EQUIPMENT_{T}_LEVEL_ALT` | `AIRBASE.{P}.equipment.{T}levelAlt`         | Plane proficiency          | `\airbaseDequipmentDlevelAlt`  | Only plane will have value       |
| `\airbase{P}equipment{T}id`       | `AIRBASE_{P}_EQUIPMENT_{T}_ID`        | `AIRBASE.{P}.equipment.{T}.data.id`         | Equipment id               | `\airbaseAequipmentAid`        |                                  |
| `\airbase{P}equipment{T}typeid`   | `AIRBASE_{P}_EQUIPMENT_{T}_TYPEID`    | `AIRBASE.{P}.equipment.{T}.data.apiTypeId`  | Equipment Type id          | `\airbaseAequipmentAtypeid`    | Diffent naming in Access         |
| `\airbase{P}equipment{T}iconid`   | `AIRBASE_{P}_EQUIPMENT_{T}_ICONID`    | `AIRBASE.{P}.equipment.{T}.data.iconTypeId` | Equipment icon id          | `\airbaseAequipmentAiconid`    | Diffent naming in Access         |
| `\airbase{P}equipment{T}equipped` | `AIRBASE_{P}_EQUIPMENT_{T}_EQUIPPED`  | ``                                          | Is Slot have Equipment     | `\airbaseAequipmentAtequipped` | No Access Method; return 0 and 1 |

- Fleet

| Mnemonic | Fleet  |
| -------- | ------ |
| A        | fleet1 |
| B        | fleet2 |
| U        | JTF    |

> `P = fleet mnemonic`

| Latex                   | Macro                   | Access                   | Usage                | Example                                                                          | Note                                                           |
| ----------------------- | ----------------------- | ------------------------ | -------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------- |
| `\fleet{P}fullAirpower` | `FLEET_{P}_AIRPOWER`    | `Fleet.{P}.fullAirpower` | Fleet air power      | `\fleetAfullAirpower` Fleet 1 air power <br> `\fleetUfullAirpower` JTF air power |                                                                |
| `\fleet{P}los{Level}`   | `FLEET_{P}_LOS_{Level}` | `Fleet.{P}.los.{L}`      | Fleet type 33 Los    | `\fleetBlosA` Fleet 1 33(1) <br> `fleetUlosC` JTF 33(3)                          | Level: `A = 1` `B = 2` `C = 3` `D = 4`                         |
| `\fleet{P}speedKanji`   | `FLEET_{P}_SPEED_KANJI` | `Fleet.{P}.fleetSpeed`   | Fleet speed in Kanji |                                                                                  |                                                                |
| `\fleet{P}speedNum`     | `FLEET_{P}_SPEED_NUM`   | ``                       | Fleet speed in Kanji |                                                                                  | No Access Method, Return Integer, Greater number, faster speed |

- Ship

| Mnemonic | Ship                          |
| -------- | ----------------------------- |
| A        | Single / JTF 1                |
| B        | Single / JTF 2                |
| C        | Single / JTF 3                |
| D        | Single / JTF 4                |
| E        | Single / JTF 5                |
| F        | Single / JTF 6                |
| G        | Single 7 (Vanguard Formation) |
| U        | JTF 7                         |
| V        | JTF 8                         |
| W        | JTF 9                         |
| X        | JTF 10                        |
| Y        | JTF 11                        |
| Z        | JTF 12                        |

> `P = ship mnemonic`

| Latex                            | Macro                              | Access                            | Usage                           | Example                        | Note             |
| -------------------------------- | ---------------------------------- | --------------------------------- | ------------------------------- | ------------------------------ | ---------------- |
| `\ship{P}nameJp`                 | `SHIP_{P}_NAME_JP`                 | `Ship.{}.date.name`               | Ship Japanese name              | `\shipAnameJp`                 |                  |
| `\ship{P}nameEn`                 | `SHIP_{P}_NAME_EN`                 | ``                                | Ship English name               | `\shipBnameEn`                 | No Access Method |
| `\ship{P}level`                  | `SHIP_{P}_LEVEL`                   | `Ship.{}.level`                   | Ship level                      | `\shipClevel`                  |                  |
| `\ship{P}fullAirPower`           | `SHIP_{P}_FULL_AIRPOWER`           | `Ship.{}.fullAirPower`            | Ship air power                  | `\shipDfullAirPower`           |                  |
| `\ship{P}displayStatusHp`        | `SHIP_{P}_DISPLAYSTATUS_HP`        | `Ship.{}.displayStatus.HP`        | Ship (Display Status) HP        | `\shipEdisplayStatusHp`        |                  |
| `\ship{P}displayStatusFirePower` | `SHIP_{P}_DISPLAYSTATUS_FIREPOWER` | `Ship.{}.displayStatus.firePower` | Ship (Display Status) FirePower | `\shipFdisplayStatusFirePower` |                  |
| `\ship{P}displayStatusArmor`     | `SHIP_{P}_DISPLAYSTATUS_ARMOR`     | `Ship.{}.displayStatus.armor`     | Ship (Display Status) Armor     | `\shipGdisplayStatusArmor`     |                  |
| `\ship{P}displayStatusTorpedo`   | `SHIP_{P}_DISPLAYSTATUS_TORPEDO`   | `Ship.{}.displayStatus.torpedo`   | Ship (Display Status) Torpedo   | `\shipUdisplayStatusTorpedo`   |                  |
| `\ship{P}displayStatusAvoid`     | `SHIP_{P}_DISPLAYSTATUS_AVOID`     | `Ship.{}.displayStatus.avoid`     | Ship (Display Status) Avoid     | `\shipUdisplayStatusAvoid`     |                  |
| `\ship{P}displayStatusAsw`       | `SHIP_{P}_DISPLAYSTATUS_ASW`       | `Ship.{}.displayStatus.asw`       | Ship (Display Status) ASW       | `\shipVdisplayStatusAsw`       |                  |
| `\ship{P}displayStatusLos`       | `SHIP_{P}_DISPLAYSTATUS_LOS`       | `Ship.{}.displayStatus.LoS`       | Ship (Display Status) LoS       | `\shipWdisplayStatusLos`       |                  |
| `\ship{P}displayStatusLuck`      | `SHIP_{P}_DISPLAYSTATUS_LUCK`      | `Ship.{}.displayStatus.luck`      | Ship (Display Status) Luck      | `\shipXdisplayStatusLuck`      |                  |
| `\ship{P}displayStatusRange`     | `SHIP_{P}_DISPLAYSTATUS_RANGE`     | `Ship.{}.displayStatus.range`     | Ship (Display Status) Range     | `\shipYdisplayStatusRange`     | Return Integer   |
| `\ship{P}displayStatusAccuracy`  | `SHIP_{P}_DISPLAYSTATUS_ACCURACY`  | `Ship.{}.displayStatus.accuracy`  | Ship (Display Status) Accuracy  | `\shipZdisplayStatusAccuracy`  |                  |

- Ship Equipment

| Mnemonic | Ship Equipment |
| -------- | -------------- |
| A        | slot 1         |
| B        | slot 2         |
| C        | slot 3         |
| D        | slot 4         |
| E        | slot 5         |
| X        | slot extra     |

> `P = ship mnemonic` `T = equipment mnemonic`

| Latex                          | Macro                              | Access                                   | Usage                      | Example                     | Note                             |
| ------------------------------ | ---------------------------------- | ---------------------------------------- | -------------------------- | --------------------------- | -------------------------------- |
| `\ship{P}equipment{T}nameJp`   | `SHIP_{P}_EQUIPMENT_{T}_NAME_JP`   | `Ship.{P}.equipment.{T}.data.name`       | Equipment Japanese name    | `\shipDequipmentAnameJp`    |                                  |
| `\ship{P}equipment{T}nameEn`   | `SHIP_{P}_EQUIPMENT_{T}_NAME_EN`   | ``                                       | Equipment English name     | `\shipEequipmentXnameEn`    | No Access Method                 |
| `\ship{P}equipment{T}remodel`  | `SHIP_{P}_EQUIPMENT_{T}_REMODEL`   | `Ship.{P}.equipment.{T}remodel`          | Equipment level of remodel | `\shipFequipmentCremodel`   |                                  |
| `\ship{P}equipment{T}levelAlt` | `SHIP_{P}_EQUIPMENT_{T}_LEVEL_ALT` | `Ship.{P}.equipment.{T}levelAlt`         | Plane proficiency          | `\shipGequipmentDlevelAlt`  | Only plane will have value       |
| `\ship{P}equipment{T}id`       | `SHIP_{P}_EQUIPMENT_{T}_ID`        | `Ship.{P}.equipment.{T}.data.id`         | Equipment id               | `\shipAequipmentAid`        |                                  |
| `\ship{P}equipment{T}typeid`   | `SHIP_{P}_EQUIPMENT_{T}_TYPEID`    | `Ship.{P}.equipment.{T}.data.apiTypeId`  | Equipment Type id          | `\shipAequipmentAtypeid`    | Diffent naming in Access         |
| `\ship{P}equipment{T}iconid`   | `SHIP_{P}_EQUIPMENT_{T}_ICONID`    | `Ship.{P}.equipment.{T}.data.iconTypeId` | Equipment icon id          | `\shipAequipmentAiconid`    | Diffent naming in Access         |
| `\ship{P}equipment{T}equipped` | `SHIP_{P}_EQUIPMENT_{T}_EQUIPPED`  | ``                                       | Is Slot have Equipment     | `\shipAequipmentAtequipped` | No Access Method; return 0 and 1 |

- Ship Slot

| Latex             | Macro               | Access              | Usage                          | Example       | Note |
| ----------------- | ------------------- | ------------------- | ------------------------------ | ------------- | ---- |
| `\ship{P}slot{T}` | `SHIP_{P}_SLOT_{T}` | `Ship.{P}.slot.{T}` | Maxium Number of Plane in Slot | `\shipAslotA` |      |

### Export Mode

> New in 0.0.6

This is the export mode of the kancolatex. 

Which will parse and output a json of processed airbase or fleet info.

You must prepare:

- An deck json

#### Example

- `kancolatex --mode export --export-type airbase -n ./example_1_fleet.json`

   This will return a airbase info json.

- `kancolatex -m export --export-type fleet -n ./example_1_fleet.json -o output.json`

   This will store fleet info into output.json

### Translate Mode

> New in 0.0.6

This is the translate mode of the kancolatex. 

Which will translate a kancolle ship or equipment.

#### Example

- `kancolatex --mode translate --translate-type ship --translate-target 加賀改二護` 

  Return `Kaga Kai Ni Go`

- `kancolatex -m translate --translate-type equipment --translate-target 15m二重測距儀改+21号電探改二+熟練射撃指揮所 `

  Return `15m Duplex Rangefinder Kai + Type 21 Radar Kai 2 + Skilled Fire Direction Center`


## Developer

```
pip install -e .[dev]
```

`.env` config

```
DEV_MODE=TRUE
```

> This project is not open public contribution because the author have no idea
> what he want.
>
> But you can fork this project and hove fun with it.

## Special Thanks

- Apple
- poky

- [noro6](https://github.com/noro6) Author of
  [制空権シミュレータ](https://noro6.github.io/kc-web/#/)

- [myangelkamikaze](https://github.com/myangelkamikaze) Contributor of
  [ElectronicObserverEN](https://github.com/ElectronicObserverEN)
- [Jebzou](https://github.com/Jebzou) Contributor of [ElectronicObserverEN]

- [KC3Kai](https://github.com/KC3Kai/KC3Kai)
