import "./EntranceChoice.css";
import { usePopup } from "../PopupContext";

export function EntranceChoice() {
  const { openPopupEditor } = usePopup();
  const { openPopupAdmin } = usePopup();

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

