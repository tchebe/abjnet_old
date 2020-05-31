# abjnet set of microservices to gather souscription data from abidjan.net partership with nsia vie assurances

A savoir:

 NBHOURTODELSUB: "nombre d'heure a attendre avant d'envoyer le mail et vider la table souscriptions qui sont a TRAITEMENT"
 DELETEPAYSAT: "cron parametre declanchant le mail d'envoi des paiements effectués puis vide la table payment"
 NBDAYTODELPRE: "nombre de jours a attendre pour vider la table des prestations qui sont a TRAITEE"
 NBDAYTOUPPRE: "nombre de jours a attendre pour mettre le statut des prestations de TRAITEMENT a TRAITEE"
 MAJSUBAT: "cron parametre declanchant la mise a jour des souscriptions de CREE a TRAITEMENT"
 MAJPREAT: "cron parametre declanchant la mise a jour des prestations de CREE a TRAITEMENT"
 TOKENEXPIRE: "nombre de minutes avant expiration du token généré"
