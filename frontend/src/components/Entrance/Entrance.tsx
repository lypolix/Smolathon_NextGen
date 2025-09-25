import { Header } from "../Header/Header";
import { useStore } from "../../backendConnection/authLogic/store/StoreContext";
import { observer } from "mobx-react-lite";
import { useState } from "react";
import "./Entrance.css";

export const Entrance = observer(() => {
  const store = useStore();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    await store.login(email, password);
  };

  return (
    <>
      <Header />
      <h1>Вход</h1>
      <div className="form">
        <input
          type="email"
          placeholder="Введите email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Введите пароль"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button onClick={handleLogin}>Войти</button>
      </div>
    </>
  );
});
