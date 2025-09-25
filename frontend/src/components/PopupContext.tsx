import { createContext, useContext, useState} from "react";
import type { ReactNode } from "react";

type PopupContextType = {
  showPopup: string;
  openPopupEditor: () => void;
  openPopupAdmin: () => void;
  closePopup: () => void;
};

const PopupContext = createContext<PopupContextType | undefined>(undefined);

export const usePopup = () => {
  const context = useContext(PopupContext);
  if (!context) throw new Error("usePopup must be used within PopupProvider");
  return context;
};

type Props = { children: ReactNode };

export const PopupProvider = ({ children }: Props) => {
  const [showPopup, setShowPopup] = useState("");

  const openPopupEditor = () => setShowPopup("editor");
  const openPopupAdmin = () => setShowPopup("admin");
  const closePopup = () => setShowPopup("");

  return (
    <PopupContext.Provider value={{ showPopup, openPopupEditor, openPopupAdmin, closePopup }}>
      {children}
    </PopupContext.Provider>
  );
};
