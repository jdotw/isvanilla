import React from "react";
import { Vendor } from "../model/types";

type Props = {
  vendors: Vendor[];
};

const OtherProducts = (props: Props) => {
  const { vendors } = props;
  return (
    <div>
      <h3>What else is there?</h3>
      <div>
        {vendors?.map((v) => (
          <div key={v.id}>
            {v.name}
            <div>
              {v.products?.map((p) => (
                <div key={p.id}>
                  {p.name} - stock level {p.stock_level}
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default OtherProducts;
