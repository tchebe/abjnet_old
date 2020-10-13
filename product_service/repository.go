package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	pb "github.com/tchebe/abjnet/product_service/proto/product"
)

type repository interface {
	Get(id string) (*pb.Product, error)
	GetAll() ([]*pb.Product, error)
	GetCotisations(police int) ([]*pb.Etat, error)
	GetAllClientProducts(*pb.Client) ([]*pb.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func newProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

//Get gets a single product by using its id
func (repo *ProductRepository) Get(id string) (*pb.Product, error) {
	var product *pb.Product
	if os.Getenv("IN_NSIA") == "no" {
		return &pb.Product{Id: "1", Name: "CAREC TEST RETRAITE"}, nil
	}
	product.Id = id
	if err := repo.db.First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

//GetAll gets all products to display
func (repo *ProductRepository) GetAll() ([]*pb.Product, error) {
	var products []*pb.Product
	type p struct {
		Id   int
		Name string
	}
	var pro p
	if os.Getenv("IN_NSIA") == "no" {
		products = append(products, &pb.Product{Id: "1", Name: "CAREC TEST RETRAITE"})
		products = append(products, &pb.Product{Id: "2", Name: "CAREC TEST EPARGNE"})
	} else {
		rows, err := repo.db.Debug().Raw("exec dbo.lstprdweblogy").Rows()
		defer rows.Close()
		if err != nil {
			log.Printf("error %v\n", err)
			return nil, err
		}
		for rows.Next() {
			log.Println(rows)
			rows.Scan(&pro.Id, &pro.Name)
			log.Printf("content of pro %v\n", pro)
			products = append(products, &pb.Product{Id: strconv.Itoa(pro.Id), Name: pro.Name})
		}
		//for _, v := range pro {

		//}

		log.Printf("content of products %v\n", products)

	}
	if len(products) > 0 {
		return products, nil
	}
	return nil, nil

}

//GetCotisations retrieves from sunshine the cotisations on this police
func (repo *ProductRepository) GetCotisations(p int) ([]*pb.Etat, error) {
	var etat []*pb.Etat
	if os.Getenv("IN_NSIA") == "no" {
		etat = append(etat, &pb.Etat{Police: strconv.Itoa(p), Datedebuteffet: "01/01/2020", Datefineffet: "01/01/2020", Libelleproduit: "test produit", Modereglement: "ESPECE", Fractionnement: "LIBRE", Numeropayeur: "001", Nompayeur: "NSIA VIE", Telephone: "010101001", Profession: "testeur", Adresse: "Testville", Datenaissance: "10/10/2000", Lieunaissance: "testville", Nomsouscripteur: "testman", Quittance: "01010101", Prime: "5000", Datecomptable: "01/03/2020", Datequittance: "01/03/2020", Etatquittance: "SOLDEE", Montantsolde: "25000", Montantimpaye: "5000", Nombresolde: "4", Nombreimpaye: "1"})
		etat = append(etat, &pb.Etat{Police: strconv.Itoa(p), Datedebuteffet: "01/01/2020", Datefineffet: "01/01/2020", Libelleproduit: "test produit", Modereglement: "ESPECE", Fractionnement: "LIBRE", Numeropayeur: "001", Nompayeur: "NSIA VIE", Telephone: "010101001", Profession: "testeur", Adresse: "Testville", Datenaissance: "10/10/2000", Lieunaissance: "testville", Nomsouscripteur: "testman", Quittance: "01010101", Prime: "5000", Datecomptable: "01/04/2020", Datequittance: "01/04/2020", Etatquittance: "IMPAYEE", Montantsolde: "25000", Montantimpaye: "5000", Nombresolde: "4", Nombreimpaye: "1"})
	} else {
		eta := struct {
			NUMERO_POLICE         string
			DATE_DEBUT_EFFET      string
			DATE_FIN_EFFET        string
			LIBELLEPRODUIT        string
			MODE_REGLEMENT        string
			FRACTIONNEMENT        string
			CIVILITE_PAYEUR       string
			NUMERO_PAYEUR         string
			NOM_PAYEUR            string
			PRENOMS_PAYEUR        string
			TELEPHONE_PAYEUR      string
			PROFESSION_PAYEUR     string
			ADRESSE_PAYEUR        string
			DATE_NAISSANCE_PAYEUR string
			LIEU_NAISSANCE_PAYEUR string
			CIVILITE_SOUSCRIPTEUR string
			NUMERO_SOUSCRIPTEUR   string
			NOM_SOUSCRIPTEUR      string
			PRENOMS               string
			TELEPHONE_CLIENT      string
			PROFESSION_CLIENT     string
			ADRESSE_CLIENT        string
			DATE_NAISSANCE_CLIENT string
			LIEU_NAISSANCE_CLIENT string
			QUITTANCE             string
			PRIME                 string
			DATE_COMPTABLE        string
			DATE_QUITTANCE        string
			PERIODE_QUITTANCE     string
			PERIODE_FIN_QUITTANCE string
			ETAT_QUITTANCE        string
			MONTANT_SOLDEES       string
			MONTANT_IMPAYES       string
			NOMBRE_SOLDEES        string
			NOMBRE_IMPAYES        string
		}{}
		rows, err := repo.db.Debug().Raw(`SELECT JPPOLIP_WNUPO NUMEROPOLICE
		, SUBSTRING(CONVERT(VARCHAR,JPPOLIP_DEFPO),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,JPPOLIP_DEFPO),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,JPPOLIP_DEFPO),1,4) AS DATE_DEBUT_EFFET
		, SUBSTRING(CONVERT(VARCHAR,JPPOLIP_DFEPO),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,JPPOLIP_DFEPO),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,JPPOLIP_DFEPO),1,4) AS DATE_FIN_EFFET
		,JPPOLIP_LIBPR1 LIBELLEPRODUIT
		,CASE WHEN [dbo].[modereg_police] (jppolip_wnupo,'g')='B' THEN 'BANCAIRE' 
			WHEN [dbo].[modereg_police] (jppolip_wnupo,'g')='C' THEN 'CHEQUE'
		WHEN [dbo].[modereg_police] (jppolip_wnupo,'g')='E' THEN 'ESPECE' 
		ELSE [dbo].[modereg_police] (jppolip_wnupo,'g') END  MODE_REGLEMENT
		
		,CASE WHEN dbo.periodicite_police(JPPOLIP_WNUPO,'I')='M' THEN 'MENSUELLE' 
			WHEN dbo.periodicite_police(JPPOLIP_WNUPO,'I')='A' THEN 'ANUUELLE'
		WHEN dbo.periodicite_police(JPPOLIP_WNUPO,'I')='S' THEN 'SEMESTRIELLE' 
		WHEN dbo.periodicite_police(JPPOLIP_WNUPO,'I')='T'  THEN 'TRIMESTRIELLE'
		ELSE dbo.periodicite_police(JPPOLIP_WNUPO,'I')END  FRACTIONNEMENT
		
		,CASE WHEN T.X.value('TITAD[1]','VARCHAR(MAX)')='M.' THEN 'MONSIEUR' 
			WHEN T.X.value('TITAD[1]','VARCHAR(MAX)')='Mle.' THEN 'MADEMOISELLE'
		WHEN T.X.value('TITAD[1]','VARCHAR(MAX)')='Mme.' THEN 'MADAME' 
		ELSE T.X.value('TITAD[1]','VARCHAR(MAX)') END  CIVILITE_PAYEUR
		,PAYEUR.JAIDENP_WNUAD NUMERO_PAYEUR
		,PAYEUR.JAIDENP_NOMAD NOM_PAYEUR
		,PAYEUR.JAIDENP_PREAD PRENOMS_PAYEUR
		,T.X.value('TELAD[1]','VARCHAR(MAX)')TELEPHONE_PAYEUR
		,T.X.value('AD0AD[1]','VARCHAR(MAX)')PROFESSION_PAYEUR
		,T.X.value('AD1AD[1]','VARCHAR(MAX)')ADRESSE_PAYEUR
		 ,SUBSTRING(CONVERT(VARCHAR,T.X.value('DNAAD[1]','VARCHAR(MAX)')),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,T.X.value('DNAAD[1]','VARCHAR(MAX)')),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,T.X.value('DNAAD[1]','VARCHAR(MAX)')),1,4)DATE_NAISSANCE_PAYEUR
		,T.X.value('LNAAD[1]','VARCHAR(MAX)')LIEU_NAISSANCE_PAYEUR
		,CASE WHEN T1.X.value('TITAD[1]','VARCHAR(MAX)')='M.' THEN 'MONSIEUR' 
			WHEN T1.X.value('TITAD[1]','VARCHAR(MAX)')='Mle.' THEN 'MADEMOISELLE'
		WHEN T1.X.value('TITAD[1]','VARCHAR(MAX)')='Mme.' THEN 'MADAME' 
		ELSE T1.X.value('TITAD[1]','VARCHAR(MAX)') END  CIVILITE_SOUSCRIPTEUR
		,SOUSC.JAIDENP_WNUAD NUMERO_SOUSCRIPTEUR
		,SOUSC.JAIDENP_NOMAD NOM_SOUSCRIPTEUR
		,SOUSC.JAIDENP_PREAD PRENOMS
		,T1.X.value('TELAD[1]','VARCHAR(MAX)')TELEPHONE_CLIENT
		,T1.X.value('AD0AD[1]','VARCHAR(MAX)')PROFESSION_CLIENT
		,T1.X.value('AD1AD[1]','VARCHAR(MAX)')ADRESSE_CLIENT
		,SUBSTRING(CONVERT(VARCHAR,T1.X.value('DNAAD[1]','VARCHAR(MAX)')),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,T1.X.value('DNAAD[1]','VARCHAR(MAX)')),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,T1.X.value('DNAAD[1]','VARCHAR(MAX)')),1,4)DATE_NAISSANCE_CLIENT
		,T1.X.value('LNAAD[1]','VARCHAR(MAX)')LIEU_NAISSANCE_CLIENT
		,A.WNUCO QUITTANCE
		,CONVERT(DECIMAL,A.MPYCO )PRIME
		, SUBSTRING(CONVERT(VARCHAR,A.DCPCO),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DCPCO),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DCPCO),1,4) AS DATE_COMPTABLE
		, SUBSTRING(CONVERT(VARCHAR,A.DDQCO),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DDQCO),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DDQCO),1,4) AS DATE_QUITTANCE
		, SUBSTRING(CONVERT(VARCHAR,A.DDPCO),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DDPCO),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DDPCO),1,4) AS PERIODE_QUITTANCE
		, SUBSTRING(CONVERT(VARCHAR,A.DFECO),7,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DFECO),5,2)+'/'+SUBSTRING(CONVERT(VARCHAR,A.DFECO),1,4) AS PERIODE_FIN_QUITTANCE
		,CASE WHEN A.INENC='S' THEN 'SOLDEES' 
			WHEN A.INENC='' AND [dbo].[modereg_police] (jppolip_wnupo,'g')='B' AND (SELECT COUNT(*) FROM NSIACIF.JASPRLP WHERE JASPRLP_DTRPV=0 AND JASPRLP_ETAPV='' AND JASPRLP_JAQUITP_WNUCO=A.WNUCO)>0 THEN 'IMPAYE EN ATTENTE RETOUR BANQUE' 
		WHEN A.INENC='' THEN 'IMPAYES'
		END  ETAT_QUITTANCE
		,(SELECT CONVERT(DECIMAL,SUM(MPYCO)) FROM NSIACIF.JAQUITP WHERE WNUPO=A.WNUPO AND INENC='S' AND MPYCO>0)MONTANT_SOLDEES
		,(SELECT CONVERT(DECIMAL,SUM(MPYCO)) FROM NSIACIF.JAQUITP WHERE WNUPO=A.WNUPO AND INENC='' AND MPYCO>0)MONTANT_IMPAYES
		,(SELECT COUNT(*) FROM NSIACIF.JAQUITP WHERE WNUPO=A.WNUPO AND INENC='S' AND MPYCO>0)NOMBRE_SOLDEES
		,(SELECT COUNT(*) FROM NSIACIF.JAQUITP WHERE WNUPO=A.WNUPO AND INENC='' AND MPYCO>0)NOMBRE_IMPAYES
		FROM NSIACIF.JAQUITP A JOIN NSIACIF.JPPOLIP B ON A.WNUPO=B.JPPOLIP_WNUPO
						JOIN (SELECT JAIDENP_WNUAD,JAIDENP_NOMAD,JAIDENP_PREAD, CONVERT(XML,FICXML)FICXML FROM NSIACIF.JAIDENP) PAYEUR ON PAYEUR.JAIDENP_WNUAD=A.WPAID
		OUTER APPLY PAYEUR.FICXML.nodes('//JAIDENP')T(X)
		JOIN (SELECT JAIDENP_WNUAD, JAIDENP_NOMAD,JAIDENP_PREAD,CONVERT(XML,FICXML)FICXML FROM NSIACIF.JAIDENP) SOUSC ON SOUSC.JAIDENP_WNUAD=A.WUCLI
		OUTER APPLY SOUSC.FICXML.nodes('//JAIDENP')T1(X)
		WHERE WNUPO= ?
		AND INENC IN ('','S')
		AND MPYCO>0
		ORDER BY DDPCO ASC
		`, p).Rows()
		defer rows.Close()
		if err != nil {
			log.Printf("error %v\n", err)
			return nil, err
		}
		for rows.Next() {
			log.Println(rows)
			rows.Scan(&eta.NUMERO_POLICE, &eta.DATE_DEBUT_EFFET, &eta.DATE_FIN_EFFET, &eta.LIBELLEPRODUIT, &eta.MODE_REGLEMENT, &eta.FRACTIONNEMENT, &eta.CIVILITE_PAYEUR, &eta.NUMERO_PAYEUR, &eta.NOM_PAYEUR, &eta.TELEPHONE_PAYEUR, &eta.PROFESSION_PAYEUR, &eta.ADRESSE_PAYEUR, &eta.DATE_NAISSANCE_PAYEUR, &eta.LIEU_NAISSANCE_PAYEUR, &eta.CIVILITE_SOUSCRIPTEUR, &eta.NUMERO_SOUSCRIPTEUR, &eta.NOM_SOUSCRIPTEUR, &eta.PRENOMS, &eta.TELEPHONE_CLIENT, &eta.PROFESSION_CLIENT, &eta.ADRESSE_CLIENT, &eta.DATE_NAISSANCE_CLIENT, &eta.LIEU_NAISSANCE_CLIENT, &eta.QUITTANCE, &eta.PRIME, &eta.DATE_COMPTABLE, &eta.DATE_QUITTANCE, &eta.PERIODE_QUITTANCE, &eta.PERIODE_FIN_QUITTANCE, &eta.ETAT_QUITTANCE, &eta.MONTANT_SOLDEES, &eta.MONTANT_IMPAYES, &eta.NOMBRE_SOLDEES, &eta.NOMBRE_IMPAYES)
			log.Printf("content of pro %v\n", eta)
			etat = append(etat, &pb.Etat{Police: eta.NUMERO_POLICE, Datedebuteffet: eta.DATE_DEBUT_EFFET, Datefineffet: eta.DATE_FIN_EFFET, Libelleproduit: eta.LIBELLEPRODUIT, Modereglement: eta.MODE_REGLEMENT, Fractionnement: eta.FRACTIONNEMENT, Numeropayeur: eta.NUMERO_PAYEUR, Nompayeur: eta.NOM_PAYEUR, Telephone: eta.TELEPHONE_CLIENT, Profession: eta.PROFESSION_CLIENT, Adresse: eta.ADRESSE_CLIENT, Datenaissance: eta.DATE_NAISSANCE_CLIENT, Lieunaissance: eta.LIEU_NAISSANCE_CLIENT, Nomsouscripteur: eta.NOM_SOUSCRIPTEUR + " " + eta.PRENOMS, Quittance: eta.QUITTANCE, Prime: eta.PRIME, Datecomptable: eta.DATE_COMPTABLE, Datequittance: eta.DATE_QUITTANCE, Etatquittance: eta.ETAT_QUITTANCE, Montantsolde: eta.MONTANT_SOLDEES, Montantimpaye: eta.MONTANT_IMPAYES, Nombresolde: eta.NOMBRE_SOLDEES, Nombreimpaye: eta.NOMBRE_IMPAYES})
		}
		//for _, v := range pro {

		//}

		log.Printf("content of products %v\n", etat)
	}
	if len(etat) > 0 {
		return etat, nil
	}
	return nil, nil
}

func (repo *ProductRepository) GetAllClientProducts(client *pb.Client) ([]*pb.Product, error) {
	var products []*pb.Product
	if os.Getenv("IN_NSIA") == "no" {
		products = append(products, &pb.Product{Id: "1", Name: "CAREC TEST RETRAITE"})
		products = append(products, &pb.Product{Id: "2", Name: "CAREC TEST EPARGNE"})
	} else {
		//select first from nsiacif.jaidenp
		// if err:=repo.db.Debug().Raw("select top 1 from nsiacif.jaidenp").Scan().Error;err!=nil{
		// 	return nil,err
		// }
		// rows, err := repo.db.Debug().Raw("exec dbo.lstprdweblogy").Rows()
		// defer rows.Close()
		// if err != nil {
		// 	log.Printf("error %v\n", err)
		// 	return nil, err
		// }
		// for rows.Next() {
		// 	log.Println(rows)
		// 	rows.Scan(&pro.Id, &pro.Name)
		// 	log.Printf("content of pro %v\n", pro)
		// 	products = append(products, &pb.Product{Id: strconv.Itoa(pro.Id), Name: pro.Name})
		// }
		// //for _, v := range pro {

		// //}

		log.Printf("content of products %v\n", products)

	}
	return products, nil

}
