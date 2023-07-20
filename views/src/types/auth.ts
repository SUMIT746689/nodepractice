
export interface AuthLogIn {
  username: string;
  password: string;
}

export interface AuthUser {
  id: number;
  username: string;
  first_name: string;
  last_name: string;
  email?: string;
  phone_number?: string;
  role: string;
}
