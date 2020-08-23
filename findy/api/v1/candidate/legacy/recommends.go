package legacy

import (
	"time"
)

type BaseModel struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JobType struct {
	BaseModel `json:",inline"`
	Type      string `json:"job_type"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Company struct {
	BaseModel `json:",inline"`
	// description: "メール配信事業/サポート（HDEブランドで開発したセキュリテイパッケージ）"
	Description string `json:"description"`
	// employee_count: null
	EmployeeCount interface{} `json:"employee_count"`
	// established_at: "1996-11-05T00:00:00.000+09:00"
	EstablishedAt time.Time `json:"established_at"`
	// github_account: ""
	GithubAccount string `json:"github_account"`
	// image_url: "https://findy-code-images.s3-ap-northeast-1.amazonaws.com/companies/00444_hennge.png"
	ImageUrl string `json:"image_url"`
	// is_agent_contract: false
	IsAgentContract bool `json:"is_agent_contract"`
	// is_contracted: true
	IsContracted bool `json:"is_contracted"`
	// is_official: true
	IsOfficial bool `json:"is_official"`
	// is_popular: false
	IsPopular bool `json:"is_popular"`
	// location: "東京都渋谷区南平台町16番28号 Daiwa渋谷スクエア"
	Location string `json:"location"`
	// name: "HENNGE株式会社"
	Name string `json:"name"`
	// policy_number: null
	PolicyNumber interface{} `json:"policy_number"`
	// president: "小椋 一宏"
	President string `json:"president"`
	// role_id: 1
	RoleId int `json:"role_id"`
	// tech_blog_rss_url: ""
	TechBlogRssUrl string `json:"tech_blog_rss_url"`
	// tech_blog_url: ""
	TechBlogUrl string `json:"tech_blog_url"`
	// url: "https://hennge.com/jp/"
	Url string `json:"url"`
	// use_zero: false
	UseZero bool `json:"use_zero"`
}

type Job struct {
	BaseModel `json:",inline"`
	// career_up_programs: "英語学習支援（オンライン英会話全額負担、セブ島流し制度）"
	CareerUpPrograms string `json:"career_up_programs"`
	// company: {id: 555, role_id: 1, name: "HENNGE株式会社", description: "メール配信事業/サポート（HDEブランドで開発したセキュリテイパッケージ）",…}
	Company Company `json:"company"`
	// company_id: 555
	CompanyId int `json:"company_id"`
	// created_at_utc: "2020-08-19T15:26:55.000Z"
	CreatedAtUTC string `json:"created_at_utc"`
	// desired_skills: "・iOS、Androidでの開発経験↵・REST APIやサーバーサイド(Python/Golangなど)での開発経験↵・AngularJS/React/VueなどのFW利用経験や知識↵・アプリケーションまたはウェブページの国際化（i18n）の実装経験"
	DesiredSkills string `json:"desired_skills"`
	// development_environment: "HTML/CSS/JavaScript"
	DevelopmentEnvironment string `json:"development_environment"`
	// employment_pattern: ""
	EmploymentPattern string `json:"employment_pattern"`
	// health: ""
	Health string `json:"health"`
	// holidays: "完全週休2日制（土・日）↵祝日、夏季↵年末年始（12/29～1/3）休暇↵有給休暇、慶弔休暇↵年間休日120日以上"
	Holidays string `json:"holidays"`
	// is_deleted: false
	IsDeleted bool `json:"is_deleted"`
	// is_liked_from_user: false
	IsLikedFromUser bool `json:"is_liked_from_user"`
	// is_liked_user: false
	IsLikedUser bool `json:"is_liked_user"`
	// is_matched: false
	IsMatched bool `json:"is_matched"`
	// is_public: true
	IsPublic bool `json:"is_public"`
	// jd_hash: "SpyQ7rV9QNezK"
	Hash string `json:"jd_hash"`
	// job_description: "HENNGEは国内シェアNo1のクラウドセキュリティサービス『HENNGE One』を主軸として、様々なソフトウェアを企業向けに提供しています。↵『HENNGE One』はOffice 365、G Suite、 Salesforce、Boxなどのクラウドサービスに対して、包括的でセキュアなサービス（SSO、アクセス制限、情報漏洩対策など）を提供しています。↵「テクノロジーの開放」をコンセプトとして、今後もコンシューマライゼーションが進んでいくであろうIoTなどの分野で、続々と新サービスを展開していく予定です。↵ ↵■フロントエンドエンジニア業務内容↵・デザイナーやバックエンドエンジニアと連携したUIの設計・実装↵・UI設計計画に基づくバックエンドAPIベースのレビューと設計↵・JavaScript/HTML/CSSを利用したWebフロントエンド開発↵・実装に必要なライブラリの構築↵・デザイナーと連携したUIコンポーネント作成 ↵↵本ポジションでは『HENNGE One』を中心としたクラウドサービスのWebフロントエンド開発を担当いただきます。↵デザイナーやバックエンドエンジニアなど各開発メンバーと密に連携し、「どのようなUI設計であれば、目的を達成できるか」を考え抜き、構築していくことがミッションとなります。フロント側からプロダクトの成長をさらに加速させる活躍を期待しています！↵ ↵またHENNGEは、エンジニアの半分以上が外国人という非常にグローバルな環境で開発を行っています。↵英語学習支援のプログラムも多数用意しており、英語を活かして仕事をしたい、今後英語を学習していきたいという方に最適な環境です。↵↵お気軽にいいかもをお待ちしております！"
	Description string `json:"job_description"`
	// job_type: {id: 1, job_type: "frontend-engineer", name: "フロントエンドエンジニア", position: 1,…}
	JobType JobType `json:"job_type"`
	// job_type_id: 1
	JobTypeId int `json:"job_type_id"`
	// location: "渋谷（東京本社）"
	Location string `json:"location"`
	// match_hash: ""
	MatchHash string `json:"match_hash"`
	// must_skills: "・HTML/CSS/JavaScriptでのSPA（シングルページアプリケーション）実装経験↵・CIサービス、GitHubなどを利用したチーム開発経験↵・英語で外国籍社員とコミュニケーションを取れる方(目安:TOEIC 650点以上) "
	MustSkills string `json:"must_skills"`
	// other_benefits: "社員持株会制度↵フリードクペ制度↵通勤手当、出張手当等支給"
	OtherBenefits string `json:"other_benefits"`
	// others: ""
	Others string `json:"others"`
	// personality: "・世の中のIT技術の動向に敏感で、初めての技術も積極的に取り入れられる方↵・課題解決に向け、主体的に考え自ら行動に移せる方 "
	Personality string `json:"personality"`
	// received_like_count: 3
	ReceivedLikeCount int `json:"received_like_count"`
	// rewarding: "・英語学習の支援プログラム多数。グローバルで戦えるエンジニアになれる絶好の環境です↵・プロダクト開発部門の担当役員が、CTOでありそしてCEOを務める、テクノロジー信仰の強い会社です↵・新しい技術や仕組みを取り入れ試していくことが会社の理念でもあるため、常に新しいスキルに触れていられます。↵・ドクターペッパー飲み放題制度あり。「優秀なエンジニアはドクペ（Dr Pepper）が好きな人が多い」というCEOの経験則から、ドクターペッパー自動販売機を社内に設置しています。 "
	Rewarding string `json:"rewarding"`
	// salary: "年収500万円〜800万円昇給　年2回賞与"
	Salary string `json:"salary"`
	// salary_examples: ""
	SalaryExample string `json:"salary_examples"`
	// salary_max: "800万円"
	SalaryMax string `json:"salary_max"`
	// salary_min: "500万円"
	SalaryMin string `json:"salary_min"`
	// score: 64
	Score int `json:"score"`
	// screening_process: ""
	ScreeningProcess string `json:"screening_process"`
	// skill_ids: [1, 5, 7, 13, 64, 73, 74, 78, 86, 87, 94, 140, 173, 199]
	SkillIds []int `json:"skill_ids"`
	// skills: [{id: 1, name: "Java"}, {id: 5, name: "Python"}, {id: 7, name: "JavaScript"}, {id: 13, name: "Go"},…]
	Skills []Skill `json:"skills"`
	// tag_ids: [1, 3, 5, 8, 11, 14, 15, 22, 23, 24, 25, 26, 28, 30, 32, 38, 43, 44]
	TagIds []int `json:"tag_ids"`
	// tags: [{id: 1, name: "社会的な影響力や意義"}, {id: 3, name: "自社サービスの開発"}, {id: 5, name: "職場の一体感"},…]
	Tags []Tag `json:"tags"`
	// title: "【グローバルな開発環境】UIの設計〜実装まで担当し、国内シェアNo1のクラウドセキュリティサービスのさらなる成長へ貢献する<フロントエンドエンジニア>を募集！ 【マザーズ上場】"
	Title string `json:"title"`
	// url: "null"
	Url string `json:"url"`
	// welfare: "雇用保険・労災保険・健康保険・厚生年金保険"
	Welfare string `json:"welfare"`
	// work_hours: "フレックスタイム制 標準勤務時間帯　10:00～19:00　　 標準労働時間 8時間"
	WorkHours string `json:"work_hours"`
	// _score: 57.600002
}

type Pagination struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_pages"`
}

type Skill struct {
	BaseModel  `json:",inline"`
	Name       string `json:"name"`
	IsOfficial bool   `json:"is_official"`
}

type Recommends struct {
	// JobTypes   []JobType  `json:"job_types"`
	JobList    []Job      `json:"list"`
	Pagination Pagination `json:"pagination"`
	// Skills     []Skill    `json:"skill"`
	Error string `json:"error"`
}
