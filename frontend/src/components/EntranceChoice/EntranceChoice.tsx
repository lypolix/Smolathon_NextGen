import "./EntranceChoice.css"

export function EntranceChoice(){
    return(
        <>
        <div className="entranceChoiceBlock">
            <button className="entranceChoiceEditor">Редактор</button>
            <button className="entranceChoiceAdmin">Администратор</button>
        </div>
        </>
    )
}