import firebase from "../firebaseConfig";
import { FirestoreDocument, BaseEntity, Display } from "../types/types"
const db = firebase.firestore();

// creates a new display and returns it
export const createDisplay = async (account_id: string, name: string): Promise<FirestoreDocument<Display>> => {
    const document = await db.collection("display").add({
        ...BaseEntity.createBaseEntity(),
        account_id,
        name
    })
    const snapshot = await document.get()
    return { id: snapshot.id, entity: snapshot.data() as Display }
}

// returns a list of all accounts that match the (should only have 1 or 0 results unless something else has gone wrong)
export const getDisplaysByAccountID = async (account_id: string): Promise<Array<FirestoreDocument<Display>>> => {
    const query_snapshot = await db.collection("display").where("account_id", "==", account_id).get();
    return query_snapshot.docs.map(s => ({
        id: s.id, entity: s.data() as Display,
    }))
};