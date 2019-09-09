package rcss

import (
	"fmt"
	"strconv"
)

type Side rune

const (
	LeftSide  Side = 'l'
	RightSide Side = 'r'
)

type UniformNumber uint8

type PlayMode string

const (
	Null            PlayMode = "null"
	BeforeKickOff   PlayMode = "before_kick_off"
	TimeOver        PlayMode = "time_over"
	PlayOn          PlayMode = "play_on"
	KickOffLeft     PlayMode = "kick_off_left"
	KickOffRight    PlayMode = "kick_off_right"
	KickInLeft      PlayMode = "kick_in_left"
	KickInRight     PlayMode = "kick_in_right"
	FreeKickLeft    PlayMode = "free_kick_left"
	FreeKickRight   PlayMode = "free_kick_right"
	CornerKickLeft  PlayMode = "corner_kick_left"
	CornerKickRight PlayMode = "corner_kick_right"
	GoalKickLeft    PlayMode = "goal_kick_left"
	GoalKickRight   PlayMode = "goal_kick_right"
	AfterGoalLeft   PlayMode = "after_goal_left"
	AfterGoalRight  PlayMode = "after_goal_right"
	DropBall        PlayMode = "drop_ball"
	OffSideLeft     PlayMode = "off_side_left"
	OffSideRight    PlayMode = "off_side_right"

	PkLeft                    PlayMode = "pk_left"
	PkRight                   PlayMode = "pk_right"
	FirstHalfOver             PlayMode = "first_half_over"
	Pause                     PlayMode = "pause"
	Human                     PlayMode = "human"
	FoulChargeLeft            PlayMode = "foul_charge_left"
	FoulChargeRight           PlayMode = "foul_charge_right"
	FoulPushLeft              PlayMode = "foul_push_left"
	FoulPushRight             PlayMode = "foul_push_right"
	FoulMultipleAttackerLeft  PlayMode = "foul_multiple_attacker_left"
	FoulMultipleAttackerRight PlayMode = "foul_multiple_attacker_right"
	FoulBallOutLeft           PlayMode = "foul_ball_out_left"
	FoulBallOutRight          PlayMode = "foul_ball_out_right"
	Max                       PlayMode = "max"
)
const (
	InitSideIndex          = 0
	InitUniformNumberindex = 1
	InitPlayMode           = 2
)

type Init struct {
	Init struct {
		Array []string `sexp:"init,siblings"`
	}
	Side          Side
	UniformNumber UniformNumber
	PlayMode      PlayMode
}

//Set Values Of init Command
func (init *Init) SetValues() {
	fmt.Println(len(init.Init.Array))
	init.Side = Side((init.Init.Array[InitSideIndex])[0])
	x, _ := strconv.ParseUint(init.Init.Array[InitUniformNumberindex], 10, 8)
	init.UniformNumber = UniformNumber(x)

	init.PlayMode = PlayMode(init.Init.Array[InitPlayMode])
	fmt.Println("ok")
}

func (m Init) adapter() *buffer {
	return &buffer{
		name: "init",
		vars: []buffer{
			{title: "side", format: "%c", value: m.Side},
			{title: "unum", value: m.UniformNumber},
			{title: "mode", value: m.PlayMode},
		},
	}
}

func (m *Init) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *Init) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
}

// Aggregate of 109 Server Parameters
type ServerParameters struct {
	ServerParameters struct {
		Array []string `sexp:"server_para,siblings"`
	}
	// Goal width
	// name: goal_width
	GoalWidth float32

	// Player size
	// name: player_size
	PlayerSize float32

	// Player decay
	// name: player_decay
	PlayerDecay float32

	//
	// name: player_rand
	PlayerRand float32

	// Player weight
	// name: player_weight
	PlayerWeight float32

	// Maximum player velocity
	// name: player_speed_max
	MaxPlayerSpeed float32

	// Maximum player acceleration
	// name: player_accel_max
	MaxPlayerAcceleration float32

	// Maximum player stamina
	// name: stamina_max
	MaxStamina float32

	// Maximum player stamina increment
	// name: stamina_inc_max
	MaxStaminaIncrement float32

	// Player recovery decrement threshold
	// name: recover_dec_thr
	PlayerRecoveryDecrementThreshold float32

	// Minimum player recovery
	// name: recover_min
	MinPlayerRecovery float32

	// Player recovery decrement
	// name: recover_dec
	PlayerRecoveryDecrement float32

	// Player dash effort decrement threshold
	// name: effort_dec_thr
	EffortDecrementThreshold float32

	// Minimum player dash effort
	// name: effort_min
	MinEffort float32

	// Player dash effort decrement
	// name: effort_dec
	EffortDecrement float32

	// Dash effort increment threshold
	// name: effort_inc_thr
	EffortIncrementThreshold float32

	// Dash effort increment
	// name: effort_inc
	EffortIncrement float32

	// Noise added directly to kicks
	// name: kick_rand
	KickRand float32

	// Flag whether to use team specific actuator noise
	// name: team_actuator_noise
	TeamActuatorNoise bool

	// Factor to multiply prand for left team
	// name: prand_factor_l
	LeftPlayerRandFactor float32

	// Factor to multiply prand for right team
	// name: prand_factor_r
	RightPlayerRandFactor float32

	// Factor to multiply kick rand for left team
	// name: kick_rand_factor_l
	LeftKickRandFactor float32

	// Factor to multiply kick rand for right team
	// name: kick_rand_factor_r
	RightKickRandFactor float32

	// Ball size
	// name: ball_size
	BallSize float32

	// Ball decay
	// name: ball_decay
	BallDecay float32

	//
	// name: ball_rand
	BallRand float32

	// Weight of the ball
	// name: ball_weight
	BallWeight float32

	// Maximum ball velocity
	// name: ball_speed_max
	MaxBallSpeed float32

	// Maximum ball acceleration
	// name: ball_accel_max
	MaxBallAcceleration float32

	// Dash power rate
	// name: dash_power_rate
	DashPowerRate float32

	//
	// name: kick_power_rate
	KickPowerRate float32

	// Kickable margin
	// name: kickable_margin
	KickableMargin float32

	// Control radius
	// name: control_radius
	ControlRadius float32

	// Goalie catch probability
	// name: catch_probability
	GoalieCatchProbability float32

	// Goalie catchable area length
	// name: catchable_area_l
	GoalieCatchableAreaLength float32

	// Goalie catchable area width
	// name: catchable_area_w
	GoalieCatchableAreaWidth float32

	// Goalie maximum moves after a catch
	// name: goalie_max_moves
	MaxGoalieAfterCatchMoves int

	// Maximum power
	// name: maxpower
	MaxPower int

	// Minumum power
	// name: minpower
	MinPower int

	// Maximum moment
	// name: maxmoment
	MaxMoment int

	// Minimum moment
	// name: minmoment
	MinMoment int

	// Maximum neck moment
	// name: maxneckmoment
	MaxNeckMoment int

	// Minimum neck moment
	// name: minneckmoment
	MinNeckMoment int

	// Maximum neck angle
	// name: maxneckang
	MaxNeckAngle int

	// Minimum neck angle
	// name: minneckang
	MinNeckAngle int

	// Visible angle
	// name: visible_angle
	VisibleAngle float32

	// Visible distance
	// name: visible_distance
	VisibleDistance float32

	// Audio cut off distance
	// name: audio_cut_dist
	AudioCutOffDistance float32

	// Quantize step of distance for movable objects
	// name: quantize_step
	MovableObjectsDistanceQuantizeStep float32

	// Quantize step of distance for landmarks
	// name: quantize_step_l
	LandmarksDistanceQuantizeStep float32

	// Quantize step of direction
	// name: quantize_step_dir
	DirectionQuantizeStep float32

	// Quantize step of distance for movable objects for left team
	// name: quantize_step_dist_team_l
	LeftTeamMovableObjectsDistanceQuantizeStep float32

	// Quantize step of distance for movable objects for right team
	// name: quantize_step_dist_team_r
	RightTeamMovableObjectsDistanceQuantizeStep float32

	// Quantize step of distance for landmarks for left team
	// name: quantize_step_dist_l_team_l
	LeftTeamLandmarksDistanceQuantizeStep float32

	// Quantize step of distance for landmarks for right team
	// name: quantize_step_dist_l_team_r
	RightTeamLandmarksDistanceQuantizeStep float32

	// Quantize step of direction for left team
	// name: quantize_step_dir_team_l
	LeftTeamDirectionQuantizeStep float32

	// Quantize step of direction for right team
	// name: quantize_step_dir_team_r
	RightTeamDirectionQuantizeStep float32

	// Corner Kick Margin
	// name: ckick_margin
	CornerKickMargin float32

	// Wind direction
	// name: wind_dir
	WindDirection float32

	//
	// name: wind_force
	WindForce float32

	//
	// name: wind_rand
	WindRand float32

	// Wind factor is none
	// name: wind_none
	NoWind bool

	// Wind factor is random
	// name: wind_random
	ProbableWind bool

	// Inertia moment for turn
	// name: inertia_moment
	InertiaMoment float32

	// Length of a half time in seconds
	// name: half_time
	// TODO: use time.Duration
	HalfTime int

	// Number of cycles to wait until dropping the ball automatically
	// name: drop_ball_time
	DropBallTime int

	// Player port number
	// name: port
	Port int

	// Offline coach port
	// name: coach_port
	OfflineCoachPort int

	// Online coach port
	// name: olcoach_port
	OnlineCoachPort int

	// Upper limit of the number of online coach’s message
	// name: say_coach_cnt_max
	OnlineCoachMaxMessageNumber int

	// Upper limit of length of online coach’s message
	// name: say_coach_msg_size
	OnlineCoachMaxMessageLength int

	// Time step of simulation [unit:msec]
	// name: simulator_step
	SimulatorStep int

	// Time step of visual information [unit:msec]
	// name: send_step
	SendStep int

	// Time step of acception of commands [unit: msec]
	// name: recv_step
	ReceiveStep int

	// Time step of body being sensed
	// name: sense_body_step
	SenseBodyStep int

	// String size of say message [unit:byte]
	// name: say_msg_size
	SayMessageSize int

	// Time window which controls how many messages can be sent (coach language)
	// name: clang_win_size
	CoachLanguageWindowSize int

	// Number of messages per window
	// name: clang_define_win
	CoachLanguageMessagesPerWindow int

	//
	// name: clang_meta_win
	CoachLanguageMetaWindow int

	//
	// name: clang_advice_win
	CoachLanguageAdviceWindow int

	//
	// name: clang_info_win
	CoachLanguageInformationWindow int

	// Delay between receipt of message and send to players
	// name: clang_mess_delay
	CoachLanguageMessageDelay int

	// Maximum number of coach messages sent per cycle
	// name: clang_mess_per_cycle
	CoachLanguageMaxMessagesPerCycle int

	//
	// name: head_max
	MaxHear int

	//
	// name: hear_inc
	HearIncrement int

	//
	// name: hear_decay
	HearDecay int

	//
	// name: catch_ban_cycle
	CatchBanCycle int

	// ?
	// name: coach
	Coach bool

	// ?
	// name: coach_w_referee
	CoachWithReferee bool

	// ?
	// name: old_coach_hear
	OldCoachHear bool

	// Interval of online coach’s look
	// name: send_vi_step
	OnlineCoachLookInterval int

	// Flag for using off side rule
	// name: use_offside
	UseOffside bool

	// Offside active area size
	// name: offside_active_area_size
	OffsideActiveAreaSize int

	// forbid kick off offside
	// name: forbid_kick_off_offside
	ForbidKickOffOffside bool

	//
	// name: log_file
	LogFile string

	//
	// name: record
	Record bool

	//
	// name: record_version
	RecordVersion int

	// Flag for record log
	// name: record_log
	RecordLog bool

	// Flag for record client command log
	// name: record_messages
	RecordMessages bool

	// Flag for send client command log
	// name: send_log
	SendLog bool

	// Flag for writing cycle lenth to log file
	// name: log_times
	LogTimes bool

	// Flag for verbose mode
	// name: verbose
	Verbose bool

	//
	// name: replay
	Replay bool

	// Offsie kick margin
	// name: offside_kick_margin
	OffsideKickMargin float32

	//
	// name: slow_down_factor
	SlowDownFactor float32

	// ?
	// name: start_goal_l
	LeftStartGoal string

	// ?
	// name: start_goal_r
	RightStartGoal string

	// ?
	// name: fullstate_l
	LeftFullState string

	// ?
	// name: fullstate_r
	RightFullState string
}

func (s *ServerParameters) SetValues() {
	// x, _ := strconv.ParseFloat(s.ServerParameters.Array[0], 64)
	// s.GoalWidth = float32(x)
	fmt.Println(len(s.ServerParameters.Array))
}
func (m ServerParameters) adapter() *buffer {
	var temp string

	return &buffer{
		name: "server_param",

		vars: []buffer{
			{name: "gwidth", value: m.GoalWidth},
			{name: "inertia_moment", value: m.InertiaMoment},
			{name: "psize", value: m.PlayerSize},
			{name: "pdecay", value: m.PlayerDecay},
			{name: "prand", value: m.PlayerRand},
			{name: "pweight", value: m.PlayerWeight},
			{name: "pspeed_max", value: m.MaxPlayerSpeed},
			{name: "paccel_max", value: m.MaxPlayerAcceleration},
			{name: "stamina_max", value: m.MaxStamina},
			{name: "stamina_inc", value: m.MaxStaminaIncrement},

			{name: "recover_init", value: temp},
			{name: "recover_dthr", value: m.PlayerRecoveryDecrementThreshold},
			{name: "recover_min", value: temp},
			{name: "recover_dec", value: m.PlayerRecoveryDecrement},
			{name: "effort_init", value: temp},
			{name: "effort_dthr", value: m.EffortDecrementThreshold},
			{name: "effort_min", value: m.MinEffort},
			{name: "effort_dec", value: m.EffortDecrement},
			{name: "effort_ithr", value: m.EffortIncrementThreshold},
			{name: "effort_inc", value: m.EffortIncrement},

			{name: "kick_rand", value: m.KickRand},
			{name: "team_actuator_noise", value: m.TeamActuatorNoise},
			{name: "prand_factor_l", value: m.LeftPlayerRandFactor},
			{name: "prand_factor_r", value: m.RightPlayerRandFactor},
			{name: "kick_rand_factor_l", value: m.LeftKickRandFactor},
			{name: "kick_rand_factor_r", value: m.RightKickRandFactor},
			{name: "bsize", value: m.BallSize},
			{name: "bdecay", value: m.BallDecay},
			{name: "brand", value: m.BallRand},
			{name: "bweight", value: m.BallWeight},

			{name: "bspeed_max", value: m.MaxBallSpeed},
			{name: "baccel_max", value: m.MaxBallAcceleration},
			{name: "dprate", value: m.DashPowerRate},
			{name: "kprate", value: m.KickPowerRate},
			{name: "kmargin", value: m.KickableMargin},
			{name: "ctlradius", value: m.ControlRadius},
			{name: "ctlradius_width", value: temp},
			{name: "maxp", value: m.MaxPower},
			{name: "minp", value: m.MinPower},
			{name: "maxm", value: m.MaxMoment},

			{name: "minm", value: m.MinMoment},
			{name: "maxnm", value: m.MaxNeckMoment},
			{name: "minnm", value: m.MinNeckMoment},
			{name: "maxn", value: m.MaxNeckAngle},
			{name: "minn", value: m.MaxNeckAngle},
			{name: "visangle", value: m.VisibleAngle},
			{name: "visdist", value: m.VisibleDistance},
			{name: "windir", value: m.WindDirection},
			{name: "winforce", value: m.WindForce},
			{name: "winang", value: temp},

			{name: "winrand", value: m.WindRand},
			{name: "kickable_area", value: temp},
			{name: "catch_area_l", value: m.GoalieCatchableAreaLength},
			{name: "catch_area_w", value: m.GoalieCatchableAreaWidth},
			{name: "catch_prob", value: m.GoalieCatchProbability},
			{name: "goalie_max_moves", value: m.MaxGoalieAfterCatchMoves},
			{name: "ckmargin", value: m.CornerKickMargin},
			{name: "offside_area", value: m.OffsideActiveAreaSize},
			{name: "win_no", value: m.NoWind},
			{name: "win_random", value: m.ProbableWind},

			{name: "say_cnt_max", value: temp},
			{name: "SayCoachMsgSize", value: temp},
			{name: "clang_win_size", value: m.CoachLanguageWindowSize},
			{name: "clang_define_win", value: temp},
			{name: "clang_meta_win", value: m.CoachLanguageMetaWindow},
			{name: "clang_advice_win", value: m.CoachLanguageAdviceWindow},
			{name: "clang_info_win", value: m.CoachLanguageInformationWindow},
			{name: "clang_mess_delay", value: m.CoachLanguageMessageDelay},
			{name: "clang_mess_per_cycle", value: m.CoachLanguageMaxMessagesPerCycle},
			{name: "half_time", value: m.HalfTime},

			{name: "sim_st", value: m.SimulatorStep},
			{name: "send_st", value: m.SendStep},
			{name: "recv_st", value: m.ReceiveStep},
			{name: "sb_step", value: m.SenseBodyStep},
			{name: "lcm_st", value: temp},
			{name: "SayMsgSize", value: m.SayMessageSize},
			{name: "hear_max", value: m.MaxHear},
			{name: "hear_inc", value: m.HearIncrement},
			{name: "hear_decay", value: m.HearDecay},
			{name: "cban_cycle", value: m.CatchBanCycle},

			{name: "slow_down_factor", value: m.SlowDownFactor},
			{name: "useoffside", value: m.UseOffside},
			{name: "kickoffoffside", value: temp},
			{name: "offside_kick_margin", value: m.OffsideKickMargin},
			{name: "audio_dist", value: m.AudioCutOffDistance},
			{name: "dist_qstep", value: m.MovableObjectsDistanceQuantizeStep},
			{name: "land_qstep", value: m.LandmarksDistanceQuantizeStep},
			{name: "dir_qstep", value: m.DirectionQuantizeStep},
			{name: "dist_qstep_l", value: m.LeftTeamMovableObjectsDistanceQuantizeStep},
			{name: "dist_qstep_r", value: m.RightTeamMovableObjectsDistanceQuantizeStep},

			{name: "land_qstep_l", value: m.LeftTeamLandmarksDistanceQuantizeStep},
			{name: "land_qstep_r", value: m.RightTeamLandmarksDistanceQuantizeStep},
			{name: "dir_qstep_l", value: m.LeftTeamDirectionQuantizeStep},
			{name: "dir_qstep_r", value: m.RightTeamDirectionQuantizeStep},
			{name: "CoachMode", value: m.Coach},
			{name: "CwRMode", value: m.CoachWithReferee},
			{name: "old_hear", value: m.OldCoachHear},
			{name: "sv_start", value: temp},
			{name: "start_goal_l", value: m.LeftStartGoal},
			{name: "start_goal_r", value: m.RightStartGoal},

			{name: "fullstate_l", value: m.LeftFullState},
			{name: "fullstate_r", value: m.RightFullState},
			{name: "drop_time", value: m.DropBallTime},
			{name: "", value: temp},
			{name: "", value: temp},
			{name: "", value: temp},
			{name: "", value: temp},
			{name: "", value: temp},
			{name: "", value: temp},
		},
	}
}

func (m *ServerParameters) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *ServerParameters) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
}

type PlayerParameters struct {
	PlayerParameters struct {
		Array []string `sexp:"player_para,siblings"`
	}
	PlayerTypes int

	SubsMax int

	PtMax int

	PlayerSpeedMaxDeltaMin   int
	PlayerSpeedMaxDeltaMax   int
	StaminaIncMaxDeltaFactor int

	PlayerDecayDeltaMin      int
	PlayerDecayDeltaMax      int
	InertiaMomentDeltaFactor int

	DashPowerRateDeltaMin int
	DashPowerRateDeltaMax int
	PlayerSizeDeltaFactor int

	KickableMarginDeltaMin int
	KickableMarginDeltaMax int
	KickRandDeltaFactor    int

	ExtraStaminaDeltaMin int
	ExtraStaminaDeltaMax int
	EffortMaxDeltaFactor int
	EffortMinDeltaFactor int

	SpareLong1  int
	SpareLong2  int
	SpareLong3  int
	SpareLong4  int
	SpareLong5  int
	SpareLong6  int
	SpareLong7  int
	SpareLong8  int
	SpareLong9  int
	SpareLong10 int

	SpareShort1  int
	SpareShort2  int
	SpareShort3  int
	SpareShort4  int
	SpareShort5  int
	SpareShort6  int
	SpareShort7  int
	SpareShort8  int
	SpareShort9  int
	SpareShort10 int
}

func (p *PlayerParameters) SetValues() {

}
func (m PlayerParameters) adapter() *buffer {
	return &buffer{
		name: "player_param",

		vars: []buffer{
			{name: "player_types", value: m.PlayerTypes},
			{name: "subs_max", value: m.SubsMax},
			{name: "pt_max", value: m.PtMax},
			{name: "player_speed_max_delta_min", value: m.PlayerSpeedMaxDeltaMin},
			{name: "player_speed_max_delta_max", value: m.PlayerSpeedMaxDeltaMax},
			{name: "stamina_inc_max_delta_factor", value: m.StaminaIncMaxDeltaFactor},
			{name: "player_decay_delta_min", value: m.PlayerDecayDeltaMin},
			{name: "player_decay_delta_max", value: m.PlayerDecayDeltaMax},
			{name: "inertia_moment_delta_factor", value: m.InertiaMomentDeltaFactor},
			{name: "dash_power_rate_delta_min", value: m.DashPowerRateDeltaMin},
			{name: "dash_power_rate_delta_max", value: m.DashPowerRateDeltaMax},
			{name: "player_size_delta_factor", value: m.PlayerSizeDeltaFactor},
			{name: "kickable_margin_delta_min", value: m.KickableMarginDeltaMin},
			{name: "kickable_margin_delta_max", value: m.KickableMarginDeltaMax},
			{name: "kick_rand_delta_factor", value: m.KickRandDeltaFactor},
			{name: "extra_stamina_delta_min", value: m.ExtraStaminaDeltaMin},
			{name: "extra_stamina_delta_max", value: m.ExtraStaminaDeltaMax},
			{name: "effort_max_delta_factor", value: m.EffortMaxDeltaFactor},
			{name: "effort_min_delta_factor", value: m.EffortMinDeltaFactor},
		},
	}
}

func (m *PlayerParameters) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *PlayerParameters) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
}

type PlayerType struct {
	PlayerType struct {
		array []string `sexp:"player_type,siblings"`
	}
	// Player Identifier
	Id int

	// Maximum Player Speed
	PlayerSpeedMax int

	// Maximum Stamina Increase
	StaminaIncMax int

	// Player Decay
	PlayerDecay int

	// Inertia Moment
	InertiaMoment int

	// Dash Power Rate
	DashPowerRate int

	// Player Size
	PlayerSize int

	// Kickable Margin
	KickableMargin int

	// Kick Rand
	KickRand int

	// Extra Stamina
	ExtraStamina int

	// Maximum Effort
	EffortMax int

	// Minimum Effort
	EffortMin int

	SpareLong1  int
	SpareLong2  int
	SpareLong3  int
	SpareLong4  int
	SpareLong5  int
	SpareLong6  int
	SpareLong7  int
	SpareLong8  int
	SpareLong9  int
	SpareLong10 int
}

func (p *PlayerType) SetValues() {

}
func (m *PlayerType) adapter(msg Message) *buffer {
	if _, err := fmt.Sscan(msg.values[0], &m.Id); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing id: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[1], &m.PlayerSpeedMax); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing player_speed_max: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[2], &m.StaminaIncMax); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing stamina_inc_max: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[3], &m.PlayerDecay); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing player_decay: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[4], &m.InertiaMoment); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing inertia_moment: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[5], &m.DashPowerRate); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing dash_power_rate: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[6], &m.PlayerSize); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing player_size: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[7], &m.KickableMargin); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing kickable_margin: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[8], &m.KickRand); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing kick_rand: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[9], &m.ExtraStamina); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing extra_stamina: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[10], &m.EffortMax); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing effort_max: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[11], &m.EffortMin); err != nil {
		return &buffer{}
		fmt.Errorf("error on parsing effort_min: %s", err)
	}
	return &buffer{}
}

func (m *PlayerType) UnmarshalRcss(msg Message) error {
	return m.adapter(msg).UnmarshalRcss(msg)
}

func (m *PlayerType) MarshalRcss(msg Message) (Message, error) {
	return m.adapter(msg).MarshalRcss()
}

// Input Driver
type Team interface {
	Name() string

	Init(match Match, side Side, unum UniformNumber, mode PlayMode)
	ServerParam(sp ServerParameters)
	PlayerParam(pp PlayerParameters)
	PlayerType(pt PlayerType)

	See()
	Hear()
	SenseBody()

	Score()
}
