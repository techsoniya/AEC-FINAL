package main

import "fmt"

func main() {
myMap := make(map[string]int)

for {
	fmt.Println("\nMenu:")
	fmt.Println("1. Create")
	fmt.Println("2. Delete")
	fmt.Println("3. Update")
	fmt.Println("4. Display")
	fmt.Println("5. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
		case 1:
			var key string
			var value int
			fmt.Print("Enter key and value: ")
			fmt.Scan(&key, &value)

			myMap[key] = value
			fmt.Println("Key-value pair created.")

		case 2:
			var keyToDelete string
			fmt.Print("Enter a key to delete: ")
			fmt.Scan(&keyToDelete)
			delete(myMap, keyToDelete)
			fmt.Println("Key deleted.")

		case 3:
			var keyToUpdate string
			fmt.Print("Enter a key to update: ")
			fmt.Scan(&keyToUpdate)
			if _, exists := myMap[keyToUpdate]; exists {
				var newValue int
				fmt.Print("Enter the new value: ")
				fmt.Scan(&newValue)
				myMap[keyToUpdate] = newValue
				fmt.Println("Value updated.")
			} else {
				fmt.Println("Key not found.")
			}

		case 4:
			fmt.Println("\nMap contents:")

			for key, value := range myMap {
				fmt.Printf("Key: %s, Value: %d\n", key, value)
			}

		case 5:
			fmt.Println("Exiting.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
