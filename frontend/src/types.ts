export interface Admin1 {
  email: string;
  role: string;
}

export interface News {
  id: number;
  title: string;
  content: string;
  tag: string;
  date: string;
  created_at: string;
  updated_at: string;
}

export interface Services {
  id: number;
  title: string;
  description: string;
  price: number;
  category: string;
  icon_url: string | null;    // может быть null в ответах
  created_at: string;
  updated_at: string;
}

export interface Team {
  id: number;
  name: string;
  position: string;
  experience: string;
  photo_url: string | null;   // может быть null в ответах
  created_at: string;
  updated_at: string;
}

export interface Projects {
  id: number;
  title: string;
  description: string;
  category: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface Statistics {
  violations_total: number;
  orders_total: number;
  fines_amount_total: number;
  collected_amount_total: number;
  evacuators_count: number;
  trips_count: number;
  evacuations_count: number;
  fine_lot_income: number;
  traffic_lights_active: number;
}

// Соответствует { "traffic": { "light_types": {...}, "install_years": {...} } }
export interface Traffic {
  light_types: Record<string, number>;
  install_years: Record<string, number>;
}
