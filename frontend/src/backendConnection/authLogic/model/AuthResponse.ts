export interface AuthResponse {
  token: string;
  user: {
    id: number;
    email: string;
    role: "admin" | "editor";
    is_active: boolean;
    created_at: string;
    updated_at: string;
  };
}