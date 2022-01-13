import "./App.css";
import axios from "axios";
import { useEffect, useState } from "react";

const apiURL = "http://localhost:8080";

function App() {
  const [isInStock, setIsInStock] = useState(false);
  const [stockLevelError, setStockLevelError] = useState(null);
  const [isLoadingStockLevel, setIsLoadingStockLevel] = useState(true);
  const [vendors, setVendors] = useState();

  useEffect(() => {
    (async () => {
      try {
        const { data: vendors } = await axios.get(`${apiURL}/vendors`);
        for (let i = 0; i < vendors.length; i++) {
          const v = vendors[i];
          const { data: products } = await axios.get(
            `${apiURL}/vendor/${v.id}/products`
          );
          v.products = products;

          if (v.name === "gloria_jeans") {
            const sfv = products.find((p) => p.name === "Sugar Free Vanilla");
            if (sfv) {
              setIsInStock(sfv.stockLevel > 0);
            } else {
              setStockLevelError(
                new Error("Failed to find Sugar Free Vanilla product")
              );
              console.error("Failed to find Sugar Free Vanilla product");
            }
          }
          vendors[i] = v;
        }
        setVendors(vendors);
        console.log("VENDORS: ", vendors);
      } catch (error) {
        setStockLevelError(error);
        console.error("EXCEPTION: ", error.error);
      } finally {
        setIsLoadingStockLevel(false);
      }
    })();
    return () => {};
  }, []);

  return (
    <div className="App">
      {isLoadingStockLevel ? (
        "Loading..."
      ) : (
        <>
          <h1>Is Sugar Free Vanilla Syrup in Stock?</h1>
          {stockLevelError ? (
            "No idea 🤷‍♂️"
          ) : (
            <h2>{isInStock ? "YES! 😃 🎉" : "No... FML. 🤦‍♂️ 😖 😫"}</h2>
          )}
          <h3>What else is there?</h3>
          <ul>
            {vendors?.map((v) => (
              <li>
                {v.name}
                <ul>
                  {v.products?.map((p) => (
                    <li>
                      {p.name} - stock level {p.stock_level}
                    </li>
                  ))}
                </ul>
              </li>
            ))}
          </ul>
        </>
      )}
    </div>
  );
}

export default App;
