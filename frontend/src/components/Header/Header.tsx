import { NavLink, useNavigate } from "react-router-dom";
import { observer } from "mobx-react-lite";
import { useState } from "react";
import { useStore } from "../../backendConnection/authLogic/store/StoreContext";
import { EntranceChoice } from "../EntranceChoice/EntranceChoice";
import "./Header.css";

type HeaderProps = {
  menuBgColor?: string;
  textColor?: string;
  activeTextColor?: string;
  buttonColor?: string;
};

export const Header = observer(
  ({
    menuBgColor = "#18211b66",
    textColor = "#F3FFED",
    activeTextColor = "#F3FFED",
    buttonColor= "#364736"
  }: HeaderProps) => {
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
              style={{ backgroundColor: buttonColor }}
              onClick={() => setShowDropdown((prev) => !prev)}
              onMouseEnter={() => setShowDropdown(true)}
              onMouseLeave={() => setShowDropdown(false)}
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

    const menuLinks = [
      { to: "/", label: "О ЦОДД", end: true },
      { to: "/news", label: "Новости" },
      { to: "/statistics", label: "Статистика" },
      { to: "/services", label: "Услуги" },
      { to: "/team", label: "Команда" },
      { to: "/projects", label: "Проекты" },
    ];

    return (
      <div className="header">
        <img className="logo" src="logo.png" alt="Logo" />
        <div className="menu" style={{ backgroundColor: menuBgColor }}>
          {menuLinks.map(({ to, label, end }) => (
            <NavLink
              key={to}
              to={to}
              end={end}
              className="menu-item"
              style={({ isActive }) => ({
                color: isActive ? activeTextColor : textColor,
              })}
            >
              {label}
            </NavLink>
          ))}
        </div>
        {renderMainButton()}
        {store.isAuth && <button style={{ backgroundColor: buttonColor }} onClick={() => store.logout()}>Выйти</button>}
      </div>
    );
  }
);
