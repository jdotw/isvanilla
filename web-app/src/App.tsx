import { useEffect, useState } from "react";
import axios from "axios";
import OtherProducts from "./components/OtherProducts";
import "./App.css";
import { Vendor, Product } from "./model/types";
import SFVStockHero from "./components/SFVStockHero";
import Loading from "./components/Loading";

const apiURL = "http://localhost:8080";

function App() {
  const [isInStock, setIsInStock] = useState(false);
  const [stockLevelError, setStockLevelError] = useState<Error>();
  const [isLoadingStockLevel, setIsLoadingStockLevel] = useState(true);
  const [vendors, setVendors] = useState<Vendor[]>();

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
            const sfv = products.find(
              (p: Product) => p.name === "Sugar Free Vanilla"
            );
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
      } catch (e) {
        setStockLevelError(e as Error);
        console.error("EXCEPTION: ", (e as Error).message);
      } finally {
        setIsLoadingStockLevel(false);
      }
    })();

    return () => {};
  }, []);

  return (
    <div className="App">
      {isLoadingStockLevel ? (
        <Loading />
      ) : (
        <>
          <SFVStockHero
            isInStock={isInStock}
            stockLevelError={stockLevelError}
          />
          {vendors && <OtherProducts vendors={vendors} />}
        </>
      )}
    </div>
  );
}

export default App;
