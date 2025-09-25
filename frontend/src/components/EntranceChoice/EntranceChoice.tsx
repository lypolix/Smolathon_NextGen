import "./EntranceChoice.css";
import { usePopup } from "../PopupContext";

export function EntranceChoice() {
  const { openPopupEditor, openPopupAdmin } = usePopup();

  return (
    <div className="entranceChoiceBlock">
      <button className="entranceChoiceEditor" onClick={openPopupEditor}>
        Редактор
      </button>
      <button className="entranceChoiceAdmin" onClick={openPopupAdmin}>
        Администратор
      </button>
    </div>
  );
}
