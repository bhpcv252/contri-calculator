# Contribution Calculator
--------
The Contribution Calculator is a simple tool designed to help groups split bills and expenses fairly.

### How to Run the Project
Execute the following command to run the project:
```
make run
```

To generate a binary file, run:
```
make build
```
The binary file will be available in the ```bin``` folder.

### Usage

- Enter the final amount which was spent? 
- Enter User's information, or leave empty to finish
  - Person's name
  - Amount paid, if any (optional)
  - Budget; how much user can afford (optional)

### Example

```
(Name Amount Budget): John 0 100
(Name Amount Budget): Jane 200
(Name Amount Budget): Dave
(Name Amount Budget): Diana 350

```
##### Output 
```
John -> 100.00 (Max Affordability)
Jane -> 100.00
Dave -> 300.00
Diana -> -50.00 (To be Collected)
```
