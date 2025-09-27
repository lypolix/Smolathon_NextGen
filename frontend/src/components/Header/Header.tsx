import { NavLink, useNavigate } from "react-router-dom";
import { observer } from "mobx-react-lite";
import { useState } from "react";
import { useStore } from "../../backendConnection/authLogic/store/StoreContext";
import { EntranceChoice } from "../EntranceChoice/EntranceChoice";
import "./Header.css";
import { usePopup } from "../PopupContext";
import AuthService from "../../backendConnection/authLogic/services/AuthService";

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
    buttonColor = "#364736",
  }: HeaderProps) => {
    const store = useStore();
    const navigate = useNavigate();
    const { showPopup, closePopup } = usePopup();
    const [showDropdown, setShowDropdown] = useState(false);

    const [adminEmail, setAdminEmail] = useState("");
    const [adminPassword, setAdminPassword] = useState("");
    const [editorEmail, setEditorEmail] = useState("");
    const [editorPassword, setEditorPassword] = useState("");

    const handleAdminLogin = async () => {
      try {
        const response = await AuthService.loginAdmin(adminEmail, adminPassword);
        localStorage.setItem("token", response.data.token);
        store.setUser(response.data.user);
        store.setAuth(true);
        closePopup();
        navigate("/admin");
      } catch (error: any) {
        alert(error.response?.data?.error || "Ошибка входа");
      }
    };

    const handleEditorLogin = async () => {
      try {
        const response = await AuthService.loginEditor(editorEmail, editorPassword);
        localStorage.setItem("token", response.data.token);
        store.setUser(response.data.user);
        store.setAuth(true);
        closePopup();
        navigate("/"); 
      } catch (error: any) {
        alert(error.response?.data?.error || "Ошибка входа");
      }
    };

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
              //onMouseLeave={() => setShowDropdown(false)}
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
      <>
        <div className="header">
          <img className="logo" src="/logo.png" alt="Logo" />
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
          {store.isAuth && (
            <button
              style={{ backgroundColor: buttonColor }}
              onClick={() => store.logout()}
            >
              Выйти
            </button>
          )}
        </div>

        {showPopup === "editor" && (
          <>
            <div className="overlay" onClick={closePopup}></div>
            <div className="mainPagePopup">
              <img onClick={closePopup} className="popupClose" src="/close.png" />
              <div className="popupHeading">
                <h2 className="popupHeadingName">Вход как</h2>
                <div className="popupHeadingName1">Редактор</div>
              </div>
              <input
                placeholder="Логин"
                className="popupInputLogin"
                value={editorEmail}
                onChange={(e) => setEditorEmail(e.target.value)}
              />
              <input
                placeholder="Пароль"
                className="popupInputPassword"
                type="password"
                value={editorPassword}
                onChange={(e) => setEditorPassword(e.target.value)}
              />
              <button className="popupButton" onClick={handleEditorLogin}>
                Войти
              </button>
            </div>
          </>
        )}

        {showPopup === "admin" && (
          <>
            <div className="overlay" onClick={closePopup}></div>
            <div className="mainPagePopup">
              <img onClick={closePopup} className="popupClose" src="/close.png" />
              <div className="popupHeading1">
                <h2 className="popupHeadingName">Вход как</h2>
                <div className="popupHeadingName11">Администратор</div>
              </div>
              <input
                placeholder="Логин"
                className="popupInputLogin"
                value={adminEmail}
                onChange={(e) => setAdminEmail(e.target.value)}
              />
              <input
                placeholder="Пароль"
                className="popupInputPassword"
                type="password"
                value={adminPassword}
                onChange={(e) => setAdminPassword(e.target.value)}
              />
              <button className="popupButton" onClick={handleAdminLogin}>
                Войти
              </button>
            </div>
          </>
        )}
      </>
    );
  }
);
