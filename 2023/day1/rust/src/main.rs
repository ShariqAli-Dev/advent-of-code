/* 
*** problem ***:
given input.txt file
each group seperate by newline are elf's calories
return the elf with the highest amount of calories

*** solution ***:
read the file

save the file into an array of characters
    - remove whitespace that may be in the saved line

have a l pointer
while l pointer is less than the length of the savedFile
    elfCalories
    while currentLine is not ""
        elfCalories += calorioes
        l += 1
    maxCal = max(elfCalories, max)
print(maxCalories)
*/
use std::fs;

fn main() {
    println!("sanith check");
    let contents = fs::read_to_string("input.txt").expect("cannot find the file");

    println!("{:?}", contents);
}
