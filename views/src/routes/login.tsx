import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { incrementByAmount } from "@/redux/features/users";
import { useAppDispatch, useAppSelector } from "@/redux/hooks";
import { useLoginUserMutation } from "@/redux/services/auth";
import { useGetAllPokemonQuery, useGetPokemonByNameQuery } from "@/redux/services/user";
import { ofetch } from "ofetch";
import { Form, redirect, useNavigate } from "react-router-dom";

export default function Login() {
  const [loginUser] = useLoginUserMutation();
  const navigate = useNavigate()

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    const username = e.target['username'].value;
    const password = e.target['password'].value;
    console.log({ username, password });
    const { data, error, loading } = await loginUser({ username, password });
    console.log({ data, error, loading });
    if (!data) return;
    if (data.env.toLowerCase() === "development") {
      document.cookie = `Authorization=${data.token}`
    }
    navigate('/dashboard')
  }

  return (
    <>
      <Form onSubmit={handleSubmit} method="post" action="/login" className="grid w-full max-w-sm items-center gap-1.5 mx-auto min-h-screen content-center">
        <Label>Username</Label>
        <Input className="max-w-sm" required name="username" />
        <Label>Password</Label>
        <Input className="max-w-sm" type="password" required name="password" />
        <Button type="submit">Login</Button>
      </Form>
    </>
  )
}

export async function LoginRouteAction({ request }) {
  //   const formData = await request.formData()
  //   if (request.method === "POST") {
  //     const { token, env } = await ofetch("/api/v1/login", { method: 'POST', baseURL: 'http://localhost:8080', body: formData }).catch(err => {
  //       return {success: false}
  //     })
  //     if (env.toLowerCase() === "development") {

  //       document.cookie = `Authorization=${token}`
  //     }
  //     return redirect("/")
  //   }
  throw new Response("", { status: 405 })
}