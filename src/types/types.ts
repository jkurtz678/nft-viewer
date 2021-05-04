import firebase from "firebase";

export interface BaseDocument {
    created_at: firebase.firestore.Timestamp;
    updated_at: firebase.firestore.Timestamp;
}

export class BaseEntity {

    static createBaseEntity(): BaseDocument {
        const now = firebase.firestore.Timestamp.now();

        return { created_at: now, updated_at: now };
    }

    static updateBaseEntity(): Partial<BaseDocument> {
        const now = firebase.firestore.Timestamp.now();

        return { updated_at: now }
    }
}

export interface FirestoreDocument<Entity extends BaseDocument> {
    entity: Entity;
    id: string;
}

export interface Account extends BaseDocument {
    address: string;
    signature: string;
    device_ids: Array<string>;
}

export interface Display extends BaseDocument {
    account_id: string;
    name: string;
    token_media_url: string;
}