import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { BrowserRouter } from "react-router-dom";
import App from "./App.tsx";
import Store from "./backendConnection/authLogic/store/store.ts";
import { StoreContext } from "./backendConnection/authLogic/store/StoreContext.ts";
const store = new Store();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <StoreContext.Provider value={store}>
        <App />
      </StoreContext.Provider>
    </BrowserRouter>
  </StrictMode>
);
