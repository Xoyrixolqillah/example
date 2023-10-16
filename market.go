package main

import (
  "database/sql"
  "fmt"
  "log"
)

type HavasStoreProduct struct {
  ID int
  ProductType string
  ProductName string
  ProductPrice float64
  ProductInformation string
  ProductQualityAssurance string
  ProductNetWeight string
  ProductValue float64
  ProductDateOfManufacture string
  ProductExpirationDate string
  ProductFactory string
  ProductDiscount float64
}

type ShoppingCart struct {
  Items []HavasStoreProduct
  Total float64
}

type CustomerLoyaltyProgram struct {
  CustomerID int
  Points int
}


func connectToDatabase() (*sql.DB, error) {
  db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/havas_store")
  if err != nil {
    return nil, err
  }
  return db, nil
}


func getAllHavasStoreProducts(db *sql.DB) ([]HavasStoreProduct, error) {
  var havasStoreProducts []HavasStoreProduct

  rows, err := db.Query("SELECT * FROM grocery_store_products")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var havasStoreProduct HavasStoreProduct

    err := rows.Scan(&havasStoreProduct.ID, &havasStoreProduct.ProductType, &havasStoreProduct.ProductName, &havasStoreProduct.ProductPrice, &havasStoreProduct.ProductInformation, &havasStoreProduct.ProductQualityAssurance, &havasStoreProduct.ProductNetWeight, &havasStoreProduct.ProductValue, &havasStoreProduct.ProductDateOfManufacture, &havasStoreProduct.ProductExpirationDate, &havasStoreProduct.ProductFactory, &havasStoreProduct.ProductDiscount)
    if err != nil {
      return nil, err
    }

    havasStoreProducts = append(havasStoreProducts, havasStoreProduct)
  }

  return havasStoreProducts, nil
}


func addToShoppingCart(shoppingCart *ShoppingCart, havasStoreProduct HavasStoreProduct) {
  shoppingCart.Items = append(shoppingCart.Items, havasStoreProduct)
  shoppingCart.Total += havasStoreProduct.ProductPrice
}


func calculateTotalPrice(shoppingCart ShoppingCart) float64 {
  var totalPrice float64

  for _, item := range shoppingCart.Items {
    totalPrice += item.ProductPrice
  }

  return totalPrice
}


func applyCustomerLoyaltyProgramDiscount(shoppingCart ShoppingCart, customerLoyaltyProgram CustomerLoyaltyProgram) float64 {
  discount := shoppingCart.Total * customerLoyaltyProgram.Points / 100

  return discount
}


func checkout(shoppingCart ShoppingCart) {
  
}

  



func main() {
	//  ma'lumotlar bazasiga ulanish
	db, err := connectToDatabase()
	if err != nil {
	  log.Fatal(err)
	}
	defer db.Close()
  
	
	havasStoreProducts, err := getAllHavasStoreProducts(db)
	if err != nil {
	  log.Fatal(err)
	}
  
	
	shoppingCart := ShoppingCart{}
  
	
	addToShoppingCart(&shoppingCart, havasStoreProducts[0])
	addToShoppingCart(&shoppingCart, havasStoreProducts[1])
  
	
	totalPrice := calculateTotalPrice(shoppingCart)
  
	
	customerLoyaltyProgram := CustomerLoyaltyProgram{
	  CustomerID: 1,
	  Points: 100,
	}
  
	discount := applyCustomerLoyaltyProgramDiscount(shoppingCart, customerLoyaltyProgram)
  
	
	checkout(shoppingCart)
  
	
	fmt.Println("Total price after discount:", totalPrice - discount)
  }
  