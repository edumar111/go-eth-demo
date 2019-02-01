pragma solidity ^0.4.24;

import "./EverisDAExpertise.sol";

contract ExpertisePopulator {

    constructor(address expertiseContract) {
        string (catC) = "courses";
        string (da, spa, ms, dops, paas, bc, bd, api) = ("Digital Architecture", "Single-Page Applications", "Microservices", "DevOps", "Platform-as-a-Service", "Blockchain", "Big Data Architectures", "API Management");
        string (b, i, a, c) = ("beginner", "intermediate", "advanced", "certified");
        EverisDAExpertise exp = EverisDAExpertise(expertiseContract);

         //////////////////////////////////
        ////// Digital Architecture //////
        addBeginner(da, "Digital Architecture Introduction", "daintro");
        addBeginner(da, "Design Thinking Principles", "desthprin");
        addBeginner(da, "Introduction to Lean", "introlean");
        addBeginner(da, "Rapid Prototyping", "rapidproto");
        addIntermediate(da, "EverisWinder Introduction", "evwinderintro");
        addIntermediate(da, "Design Thinking: 90-minute Introduction", "desth90mintro")
        addIntermediate(da, "Reactive Spring and Spring Boot", "reactspringboot");
        addAdvanced(da, "EverisWinder, Advanced Topics", "evwinderadv");
        addAdvanced(da, "Architecture by Example", "archbyex");
        addCertified(da, "Oracle Java Architect Certification");

        ////// SPA //////
        addBeginner(spa, "Lean Architectures Introduction", "leanarchintro");
        addBeginner(spa, "Angular 2 Introduction", "angular2intro");
        addBeginner(spa, "AngularJS Basics", "angularjsbasics");
        addBeginner(spa, "Angular 2 Basics", "angular2basics");
        addIntermediate(spa, "Spring: Web Applications with Spring and Angular", "springangwebapps");
        addIntermediate(spa, "A Practical Intro to React.js", "practicalintroreactjs");
        addAdvanced(spa, "SPA State of the art", "spastateoftheart");
        addAdvanced(spa, "Developing Angular web apps with Typescript", "devangwebappstypescript");

        ////// Microservices //////
        addBeginner(ms, "Microservices Introduction", "msintro");
        addBeginner(ms, "Spring Cloud Microservices", "springcloudms");
        addBeginner(ms, "Microservices Workshop", "msworkshop");
        addBeginner(ms, "From Monolith to Microservices", "frommonolithtoms");
        addIntermediate(ms, "Building Microservices with Spring Boot, Spring Cloud and Cloud Foundry", "msspringbootcloudfoundry");
        addIntermediate(ms, "Microservices Architecture and Design", "msarchdesign");
        addAdvanced(ms, "Scalable Microservices with Kubernetes", "scalablemskub");
        addAdvanced(ms, "Openshift Microservices Building and Execution", "openshiftmsbuildexec");
        addAdvanced(ms, "Microservices with Java, Scala and Akka", "msjavascalaakka");

        ////// DevOps //////
        addBeginner(dops, "What is DevOps?", "whatisdops");
        addBeginner(dops, "DevOps Introduction", "dopsintro");
        addBeginner(dops, "DevOps Architecture", "dopsarch");
        addBeginner(dops, "DevOps Alimentatec", "dopsalimentatec");
        addBeginner(dops, "The DevOps Toolkit", "dopstoolkit");
        addBeginner(dops, "Getting Started with DevOps in 90 Minutes", "gettingstarteddevops90m");
        addIntermediate(dops, "Devops with Altemista Platform", "dopsaltemista");
        addIntermediate(dops, "Jenkins Basics", "jenkinsbasics");
        addIntermediate(dops, "Containers Workshop", "contworkshop");
        addIntermediate(dops, "Continuous Integration – Jenkins", "cijenkins");
        addIntermediate(dops, "Quality Assurance – Selenium and Cucumber", "qaseleniumcucumber");
        addIntermediate(dops, "Artifact Management – Nexus", "artifactnexus");
        addIntermediate(dops, "Version Control System – Git & Gerrit", "gitgerrit");
        addIntermediate(dops, "Team Collaboration – Jira, Confluence, Hipchat", "jiraconfluencehipchat");
        addIntermediate(dops, "Load Testing Framework – Gatling Static Code Analysis – SonarCube", "loadtestgatlingsonarcube");
        addIntermediate(dops, "Containers – Docker", "contdocker");
        addAdvanced(dops, "Devops & Design Thinking", "dopsdesth");
        addAdvanced(dops, "Advanced Jenkins", "advjenkins");
        addAdvanced(dops, "DevOps Workshop", "dopsworkshop");
        addAdvanced(dops, "Monitoring – Sensu, Uchiwa Infra Provisioning & Conf Automation – Ansible, Terraform", "monsensuuchiwaansibletf");
        addAdvanced(dops, "Continuous Security – OWASP", "contsecowasp");
        addCertified(dops, "Puppet Certification");

        ////// PaaS //////
        addBeginner(paas, "Introducción a Arquitecturas Cloud", "introarqucloud");
        addBeginner(paas, "Curso de Introducción al Cloud", "cursointrocloud");
        addBeginner(paas, "Cloud aPaaS", "cloudapaas");
        addIntermediate(paas, "Curso de Cloud Intermedio", "cursocloudintermedio");
        addIntermediate(paas, "Introducción a Bluemix", "introbluemix");
        i PaaS Value Proposition paasvalueproposition
        i Paths to Cloud pathstocloud
        i Introduction to Google Cloud Platform introgooglecloud
        A Taller Ejecución IaaS para Oracle Cloud Platform talleriaasoracle
        A Taller Desarrollo para Oracle Cloud Platform tallerdevoraclecloud
        A Taller Integración para Oracle Cloud Platform "tallerintegracionoracle"
        A "Analytics Cloud Services en Oracle Cloud Platform" "analyticsoraclecloud"
        A "Cloud Native Architecture Patterns" "cloudnativearchpatterns"
        C "Certificación como desarrollador PCF" "pcfdevcert"
        C "Amazon Web Services: Architect Associate Certification -AWS Core Architecture Concepts" "awsarchassociatecore"
        C "Certificación Plataforma Google Cloud" "googlecloudcert"
        C "Bluemix Certification" "bluemixcert"
        C "MS Azure Certification" "azurecert"
        C "OpenShift Certification" "openshiftcert"

        ////// Blockchain //////
        b "Introduction to Blockchain" "blockchainintro"

        "blktech","courses","Blockchain Technologies","beginner","blockchain",10
        "techintroblk","courses","Technical Introduction to Blockchain","beginner","blockchain",10
        "ripplepayments","courses","Ripple for International Payments","beginner","blockchain",10
        "smartcontoverv","courses","Smart Contracts Overview","beginner","blockchain",10

        "smartcontr","courses","Smart Contracts","intermediate","blockchain",20
        "blkworkshopartone","courses","Blockchain Workshop - Part 1","intermediate","blockchain",20
        "blkworkshopartwo","courses","Blockchain Workshop - Part 2","intermediate","blockchain",20
        "blkworkshoparthree","courses","Blockchain Workshop - Part 3","intermediate","blockchain",20

        "everisripple","courses","Everis Ripple","advanced","blockchain",30
        "blkcapmarktf","courses","Blockchain in Capital Markets and Trade Finance","advanced","blockchain",30
        "platformcompare","courses","Platform Comparision","advanced","blockchain",30
        "hyperoverview","courses","ILP & Hyperledger QuILt Overview","advanced","blockchain",30



         ////////////////////////////////
        ////// Hands-On Experience //////
        exp.createGoal("proposal-staff", "handOnExp", "Proposal (participant)", "participant", "", 30);
        exp.createGoal("proposal-lead", "handOnExp", "Proposal (lead)", "lead", "", 50);
        exp.createGoal("poc-staff", "handOnExp", "Proof of Concept (participant)", "participant", "", 40);
        exp.createGoal("poc-lead", "handOnExp", "Proof of Concept (lead)", "lead", "", 60);
        exp.createGoal("proj-staff", "handOnExp", "Project (participant)", "participant", "", 75);
        exp.createGoal("proj-lead", "handOnExp", "Project (lead)", "lead", "", 150);
        exp.createGoal("bigproj-staff", "handOnExp", "Large Project (participant)", "participant", "", 125);
        exp.createGoal("bigproj-lead", "handOnExp", "Large Project (lead)", "lead", "", 250);

         ////////////////////////
        ////// Evangelism //////
        exp.createGoal("article", "evangelism", "Article Writing", "", "", 50);
        exp.createGoal("interview", "evangelism", "Interview", "", "", 75);
        exp.createGoal("community", "evangelism", "Community Engagement", "", "", 100);
        exp.createGoal("teaching", "evangelism", "Course Teaching", "", "", 125);
        exp.createGoal("webinarteaching", "evangelism", "Webinar Teaching", "", "", 150);
        exp.createGoal("confspeak", "evangelism", "Conference Speaking", "", "", 175);
        exp.createGoal("book", "evangelism", "Book Publishing", "", "", 500);

         //////////////////////////////////
        ////// Education & Research //////
        exp.createGoal("csbach", "education", "Computing Science Bachelor", "", "", 100);
        exp.createGoal("oncampusarchcert", "education", "On-Campus Architecture Certification", "", "", 150);
        exp.createGoal("csmaster", "education", "Computing Science Master Degree", "", "", 200);
        exp.createGoal("csphd", "education", "Computing Science PhD", "", "", 400);
        exp.createGoal("archpaper", "research", "Architecture Paper Publishing", "", "", 100);
        exp.createGoal("archpatent", "research", "Architecture-related Patent", "", "", 400);

    }

    function addBeginner(string subject, string name, string id) {
        exp.createGoal(id, "courses", name, "beginner", subject, 10);
    }
    function addIntermediate(string subject, string name, string id) {
        exp.createGoal(id, "courses", name, "intermediate", subject, 20);
    }
    function addAdvanced(string subject, string name, string id) {
        exp.createGoal(id, "courses", name, "advanced", subject, 30);
    }
    function addCertified(string subject, string name, string id) {
        exp.createGoal(id, "courses", name, "certified", subject, 50);
    }
}
