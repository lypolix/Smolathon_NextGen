import { NavLink, useNavigate } from "react-router-dom";
import { observer } from "mobx-react-lite";
import "./Header.css";
import { useState } from "react";
import { useStore } from "../../authLogic/store/StoreContext";
import { EntranceChoice } from "../EntranceChoice/EntranceChoice";

export const Header = observer(() => {
  const store = useStore();
  const navigate = useNavigate();
  const [showDropdown, setShowDropdown] = useState(false);

  const renderMainButton = () => {
    if (store.user?.role === "editor") {
      return <button>Редактировать</button>;
    }
    if (store.user?.role === "admin") {
      return <button onClick={() => navigate("/admin")}>Админ-панель</button>;
    }
    if (!store.isAuth) {
      return (
        <div className="dropdown-wrapper">
          <button
            className="enterH"
            onClick={() => setShowDropdown((prev) => !prev)}
            onMouseEnter={() => setShowDropdown(true)} // появится при наведении
            onMouseLeave={() => setShowDropdown(false)} // скроется, если уйти мышкой
          >
            Вход
          </button>
          {showDropdown && (
            <div
              className="dropdown"
              onMouseEnter={() => setShowDropdown(true)}
              onMouseLeave={() => setShowDropdown(false)}
            >
              <EntranceChoice />
            </div>
          )}
        </div>
      );
    }
  };

  return (
    <div className="header">
      <img className="logo" src="logo.png" alt="Logo" />
      <div className="menu">
        <NavLink to="/" end className="menu-item">
          О ЦОДД
        </NavLink>
        <NavLink to="/news" className="menu-item">
          Новости
        </NavLink>
        <NavLink to="/projects" className="menu-item">
          Проекты
        </NavLink>
        <NavLink to="/reestr" className="menu-item">
          Реестр
        </NavLink>
        <NavLink to="/statistics" className="menu-item">
          Статистика
        </NavLink>
        <NavLink to="/team" className="menu-item">
          Команда
        </NavLink>
      </div>
      {renderMainButton()}
      {store.isAuth && <button onClick={() => store.logout()}>Выйти</button>}
    </div>
  );
});
