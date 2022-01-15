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
    </div>
  );
};

export default OtherProducts;
