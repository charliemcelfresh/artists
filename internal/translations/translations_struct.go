package translations
type Translations struct {
SqlDrivers struct {
	Postgres string `yaml:"postgres"`
} `yaml:"sql_drivers"`
Errors struct {
	ErrorLoadingEnv string `yaml:"error_loading_env"`
	ErrorConnectingToDb string `yaml:"error_connecting_to_db"`
	ErrorPerformingQuery string `yaml:"error_performing_query"`
	ErrorSubcommandRequired string `yaml:"error_subcommand_required"`
} `yaml:"errors"`
Paths struct {
	Artists string `yaml:"artists"`
} `yaml:"paths"`
Ports struct {
	EightyEighty string `yaml:"eighty_eighty"`
} `yaml:"ports"`
Flags struct {
LoadArtistsCountFlag struct {
	Description string `yaml:"description"`
	Name string `yaml:"name"`
} `yaml:"load_artists_count_flag"`
} `yaml:"flags"`
Commands struct {
Root struct {
	Use string `yaml:"use"`
	Short string `yaml:"short"`
} `yaml:"root"`
Server struct {
	Use string `yaml:"use"`
	Short string `yaml:"short"`
} `yaml:"server"`
LoadArtists struct {
	Use string `yaml:"use"`
	Short string `yaml:"short"`
} `yaml:"load_artists"`
BuildYamlStruct struct {
	Use string `yaml:"use"`
	Short string `yaml:"short"`
} `yaml:"build_yaml_struct"`
} `yaml:"commands"`
EnvironmentVariables struct {
	DatabaseUrl string `yaml:"database_url"`
} `yaml:"environment_variables"`
}
