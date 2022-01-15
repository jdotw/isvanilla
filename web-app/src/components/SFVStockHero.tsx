import React from "react";

type Props = {
  isInStock: boolean;
  stockLevelError?: Error;
};

function SFVStockHero(props: Props) {
  const { isInStock, stockLevelError } = props;
  return (
    <div>
      <h1>Is Sugar Free Vanilla Syrup in Stock?</h1>

      {stockLevelError ? (
        "No idea 🤷‍♂️"
      ) : (
        <h2>{isInStock ? "YES! 😃 🎉" : "No... FML. 🤦‍♂️ 😖 😫"}</h2>
      )}
    </div>
  );
}

export default SFVStockHero;
