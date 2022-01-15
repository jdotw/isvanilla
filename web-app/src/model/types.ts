export type Product = {
  id: string;
  name: string;
  stock_level: number;
};

export type Vendor = {
  id: string;
  name: string;
  products?: Product[];
};
