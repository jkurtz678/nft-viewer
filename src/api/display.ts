import firebase from "../firebaseConfig";
import { FirestoreDocument, BaseEntity, Display } from "../types/types"

const db = firebase.firestore();

// creates a new display and returns it
export const createDisplay = async (account_id: string, name: string, code: string): Promise<FirestoreDocument<Display>> => {
    const document = await db.collection("display").add({
        ...BaseEntity.createBaseEntity(),
        account_id,
        name,
        code
    })
    const snapshot = await document.get()
    return { id: snapshot.id, entity: snapshot.data() as Display }
}

export const createDisplayWithListener = async (account_id: string, name: string, code: string, onChange: (doc: FirestoreDocument<Display>) => void) => {
    const document = await db.collection("display").add({
        ...BaseEntity.createBaseEntity(),
        account_id,
        name,
        code
    })
    document.onSnapshot((doc) => {
        onChange({ id: doc.id, entity: doc.data() as Display })
    })
}

// returns a list of all accounts that match the (should only have 1 or 0 results unless something else has gone wrong)
export const getDisplaysByAccountID = async (account_id: string): Promise<Array<FirestoreDocument<Display>>> => {
    const query_snapshot = await db.collection("display").where("account_id", "==", account_id).get();
    return query_snapshot.docs.map(s => ({
        id: s.id, entity: s.data() as Display,
    }))
};

export const getDisplaysByAccountIDWithListener = async (account_id: string, onChange: (arr: Array<FirestoreDocument<Display>>) => void) => {
    await db.collection("display").where("account_id", "==", account_id).onSnapshot(snapshot => {
        onChange(snapshot.docs.map( d => ({id: d.id, entity: d.data() as Display})))
    })
};

export const getDisplayByDisplayID = async (display_id: string): Promise<FirestoreDocument<Display>> => {
    const document_snapshot = await db.collection("display").doc(display_id).get()
    return { id: document_snapshot.id, entity: document_snapshot.data() as Display };
}

export const getDisplayByDisplayIDWithListener = async (display_id: string, onChange: (doc: FirestoreDocument<Display>) => void) => {
    db.collection("display").doc(display_id).onSnapshot((doc) => {
        onChange({ id: doc.id, entity: doc.data() as Display });
    })
}

export const updateDisplay = async (display: FirestoreDocument<Display>): Promise<FirestoreDocument<Display>> => {
    await db.collection("display").doc(display.id).update(display.entity)
    const snapshot = await db.collection("display").doc(display.id).get()
    return { id: snapshot.id, entity: snapshot.data() as Display }
}

export const addAccountToDisplay = async(display_id: string, account_id: string): Promise<FirestoreDocument<Display>> => {
    await db.collection("display").doc(display_id).update({account_id})
    const snapshot = await db.collection("display").doc(display_id).get()
    return { id: snapshot.id, entity: snapshot.data() as Display }
}