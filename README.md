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
  - Initial contribution, if any (optional)
  - How much user can afford (optional)

### Example

```
Enter the final amount which was spent? : 1000
```

```
(Name Amount Afford): John 0 100
(Name Amount Afford): Jane 200
(Name Amount Afford): Dave
(Name Amount Afford): Diana 350

```
##### Output 
```
John -> 100.00 (Max Affordability)
Jane -> 100.00
Dave -> 300.00
Diana -> -50.00 (To be Collected)
```